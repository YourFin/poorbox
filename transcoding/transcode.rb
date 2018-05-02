#!/bin/ruby
require 'open3'

filename = ARGV[0]
filename_no_ext = filename.gsub(/\.[a-zA-Z0-9]+$/, '')
out_name = filename_no_ext + ".webm"

# get stream data:
# see http://web.archive.org/web/20180501034714/http://blog.honeybadger.io/capturing-stdout-stderr-from-shell-commands-via-ruby/
raw_stream_data, _, status = Open3.capture3("ffprobe -i #{filename} -show_streams")
## Dump ffmpeg output:
#print "\n\n"
#
#print raw_stream_data
#
#print "\n\n"

# Parse stream data into a map
counter = -1
streams = []
for line in raw_stream_data.split("\n")
  if line == '[STREAM]' then
    counter += 1
    streams[counter] = {}
  elsif line != '[\STREAM]'
    sp = line.split("=", 2)
    streams[counter][sp[0]] = sp[1]
  end 
end

# Parse maps
used_streams = []
for stream in streams
  index = stream["index"].to_i
  if index == nil then
    raise "No index for stream: #{stream}"
  end
  case stream["codec_type"]
  when "video"
    used_streams[index] = {:type => :video}
  when "audio"
    # some movies have weird sub tracks that are
    # quite needed even for english, see "Snowpiercer",
    # so we're not going to attempt to pick and
    # choose audio tracks
    if streams.select { |elem| elem["TAG:language"] == stream["TAG:language"] }.
         map { |elem| elem["channels"].to_i }.max > stream["channels"].to_i then
    else 
      used_streams[index] = {:type => :audio,
                             :channels => stream["channels"].to_i,
                             :lang => stream["TAG:language"]}
    end
  when "subtitle"
    #subtitles are tiny and should all be kept
    used_streams[index] = {:type => :sub}
    # Dump any other crap in the container file
  end
end

puts "kept streams: #{used_streams}"

# Build up ffmpeg commands
# http://web.archive.org/web/20180501174740/https://trac.ffmpeg.org/wiki/Map
maps = ""
copies = []
videoNum = -1
audioNum = -1
subNum   = -1

def audioDefaultOptions(anum)
  "-c:a:#{anum} libopus -filter:a:#{anum} loudnorm -af:a:#{anum} aformat=channel_layouts=\"7.1|5.1|stereo\" -b:a:#{anum} 64k" 
end

used_streams.each_with_index do |stream, ii| 
  maps += " -map 0:#{ii}"
  case stream[:type]
  when :video
    videoNum += 1
    copies.push "-c:v:#{videoNum} libvpx-vp9 -threads 8 -speed 4 -b:v:#{videoNum} 1000k"
  when :audio
    if stream[:channels] > 2
      audioNum += 1
      maps += " -map 0:#{ii}"
      copies.push audioDefaultOptions(audioNum) + " -ac:a:#{audioNum} 2"
    end
    audioNum += 1
    copies.push audioDefaultOptions(audioNum)
  when :sub
    subNum += 1
    copies.push "-c:s:#{subNum} webvtt"
  end
end

FFMPEG_OPTIONS = "-i '#{filename}' -f webm"

puts "ffmpeg #{FFMPEG_OPTIONS} -y -pass 1 #{maps} #{copies.join(" ")} /dev/null"
puts "\n"
puts "ffmpeg #{FFMPEG_OPTIONS} -pass 2 #{maps} #{copies.join(" ")} #{out_name}"
#Example 1:
#ffmpeg -i $1 -c:v libvpx-vp9 -b:v 2M -pass 1 -c:a libopus -f webm /dev/null && \
#    ffmpeg -i $1 -c:v libvpx-vp9 -b:v 2M -pass 2 -c:a libopus $(echo $1 | sed -E 's/\.[a-zA-Z0-9]+$//').webm

#Example 2:
#ffmpeg -y -i $1 -c:v libvpx-vp9 -pass 1 -b:v 1000K -threads 8 -speed 4 \
#-tile-columns 6 -frame-parallel 1 \
#-an -f webm /dev/null
#ffmpeg -i $1 -c:v libvpx-vp9 -pass 2 -b:v 2M -threads 8 -speed 1 \
#		    -tile-columns 6 -frame-parallel 1 -auto-alt-ref 1 \
#		    -c:a libopus -f webm \
#		    $(echo $1 | sed -E 's/\.[a-zA-Z0-9]+$//').webm
