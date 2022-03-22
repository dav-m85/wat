package internal

//	if n.Name() == ".git" {
//		err := handleGit(dir, &e)
//		if err != nil {
//			errc <- fmt.Errorf("cannot get git info in %s: %w", dir, err)
//		}
//		continue
//	}
// func handleGit(dir string, e *Entry) error {
// 	gc := exec.Command("/usr/bin/git", `for-each-ref --format="%(refname:short) %(upstream:short)" refs/heads`)
// 	gc.Dir = dir
// 	out, err := gc.Output()
// 	if err != nil {
// 		return err
// 	}
// 	type s struct {
// 		local, remote string
// 	}
// 	var br []s
// 	for _, l := range strings.Split(string(out), "\n") {
// 		x := strings.Split(l, " ")
// 		br = append(br, s{
// 			local:  x[0],
// 			remote: x[1],
// 		})
// 	}

// 	// gc := exec.Command("git", "status", "-sb")
// 	// gc.Dir = dir
// 	// out, err := gc.Output()
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	e.Git = strings.Split(string(out), "\n")[0]

// 	return nil
// }
