## Week 1

### Monday, June 4th - Wednesday, June 6th
- Orientation, tour, meet with team, etc.
- Setup environment

### Thursday, June 7th, 2018
TIL: 
- `dep` is a tool for managing dependencies for Go projects, similar to how webpack bundles packages
- to import dep, input `go get -u github.com/golang/dep/...`
- gherklin (https://github.com/cucumber/cucumber/wiki/Gherkin)
- goconvey (https://github.com/smartystreets/goconvey)

TCIL (Terminal Commands I Learned):  
1. `dep init` (set up a new project)
2. `dep ensure` (install project's dependencies)
3. `go test -v`  (list the test cases like rspec)
4. `go format <insert filename>` (standard for formatting Go code)
5. `gdir` is an alias to `~/projects/apps/gocode`

Types:
- **int**, int8, int16, int32, **int64**
- uint, uint8, uint16, uint32, uint64 (only holds positive numbers)
- float32, float64

Did you know?
Go knows what the types of your values are at compile time, and can tell whether they're being used appropriately.