package tmdb

func AddDir(path string, ignoreDotfiles bool) {
	files, _ := filepath.Glob(path)
	for _, f := range files {
		// Ignore dotfiles
		if ignoreDotfiles && f[0] == "."[0] {
			continue
		}
		//drop file extensions for search
		extnSplit := strings.Split(filepath.Base(f), ".")
		fmt.Println(extnSplit)
		var extentionless string
		if len(extnSplit) < 2 {
			extentionless = strings.Join(extnSplit, "")
		} else {
			extentionless = strings.Join(extnSplit[:len(extnSplit)-1], "")
		}
		fmt.Println(extentionless)
		// Add to db
	}
}

func getUserQuery() string {
	var query string
	fmt.Println("Please enter a new query:\n")
	fmt.Scanln(&query)
	return query[:len(query) - 1]
}

func userConfirmf(format string, a ...interface{}) bool {
	fmt.Printf("%s [y/n]: ", fmt.Sprintf(format, a...))
	var response string
	for {
		fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		// TODO: make this shorter
		switch response {
		case "y":
			return true
		case "yes":
			return true
		case "no":
			return false
		case "n":
			return false
		case "true":
			return true
		case "false":
			return false
		}
		fmt.Println("Invalid input. \"y\" or \"n\" are acceptable.\nTry again. [y/n]: ")
	}
}

func CmdMovieSearch(query string) {
	for {
		fmt.Printf("Finding results for %q...\n", query)
		queryResults, err := movieSearch(query)
		switch len(queryResults) {
		case 0:
			fmt.Println("No result found for", query)
			query = getUserQuery()
		case 1:
			if userConfirmf("Found 1 result: %q\nAccept?", queryResults[0]) {
				//return result
			} else if userConfirmf("Enter new query?") {
				query = getUserQuery()
			} else {
				// error out
			}
		default:
			for index, value := range queryResults {
				fmt.Printf("%d: %q\n", index + 1, value)
			}
			fmt.Println("0: Different query")
			fmt.Println("-1: Abort this file")
			//User input loop
			err := errors.New("No error found, but err not set to nil")
			var choice int
			for err != nil {
				fmt.Println("\nEnter a number: ")
				_, err := fmt.Scanf("%d", &choice)
				if err != nil {
					fmt.Println("Error: %q", err)
				}
				switch {
				case choice == 0:
					query = getUserQuery()
					break // breaks for loop
				case choice >= 0 && choice < len(queryResults):
					//return queryresults[choice]
					return
				default:
					// abort and error out
					return
				}
			}
		}
	}
}
