# wat
Print an excerpt of READMEs contained in current directory or its children. Useful when you lost track of what's what in your project's directory.

## Usage
Run:

    wat

If wat finds a:
- **README.md**, it prints first non-heading markdown paragraph
- **README**, it prints the first 2 lines
- ***.tar** file, it opens it and looks for a README, then apply above rules

For more options, see `wat -help`.

## Installation
So far, there's no binary release, you should have golang installed and run:

    go install github.com/dav-m85/wat

## TODO
- Detect if project is versionned with git and report status
- Explore inside gzip tar archives
- Collect more info on project directory, will require copying fs.WalkDir code

## Bibliography
- https://stackoverflow.com/questions/7773939/show-git-ahead-and-behind-info-for-all-branches-including-remotes
