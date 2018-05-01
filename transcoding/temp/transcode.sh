#!/bin/ruby



# get stream data:
raw_stream_data = `ffprobe -i #{ARGV[0]} -show_streams`
## Dump ffmpeg output:
#print "\n\n"
#
#print raw_stream_data
#
#print "\n\n"

counter = -1
streams = []

# Parse stream data into a map
for line in raw_stream_data.split("\n")
  if line == '[STREAM]' then
    counter += 1
    streams[counter] = {}
  elsif line != '[\STREAM]'
    sp = line.split("=", 2)
    streams[counter][sp[0]] = sp[1]
  end 
end

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
    used_streams[index] = {:type => :audio,
                           :channels => stream["channels"].to_i }
  when "subtitle"
    #subtitles are tiny and should all be kept
    used_streams[index] = {:type => :sub}
    # Dump any other crap in the container file
  end
  puts stream["index"] + " " + stream["codec_type"]
end

puts "kept streams: #{used_streams}"

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
