workflow "build and test" {
  on = "push"
  resolves = ["build"]
}

action "lint" {
  uses = "docker://golangci/golangci-lint"
  args = "golangci-lint run ./..."
}

action "test" {
  uses = "docker://golang:latest"
  args = "go test -v ./..."
}

action "build" {
  needs = ["lint", "test"]
  uses = "docker://golang:latest"
  args = "go build ./..."
}
