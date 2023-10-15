# What are Go modules?

Introduced in v1.1, finalized in v1.3, they let you initialize a Go project anywhere

To initialize modules: `go mod init project/package-name`
To import new dependencies: `go get package-name`

One can reference these imported packages in their go files

## Dependency management

### Dependency Requirements

What is the dependency and where to find them

```text
HostOfDependency/UserOrOrganization/Project
github.com/Suraj-Thotakura/Go-Project
```

### Dependency Versioning

Which version of the dependency to use

---

Dependency management is done with the help of two files

### go.mod

Hosts the dependency requirement and versioning in a user friendly format
Similar to package.json

### go.sum

Host the direct paths to the dependencies and the paths of the dependencies's dependencies, their versions and the check sums (to ensure the integrity of these dependencies)
Similar to package-lock.json

---

## Other commands

For adding missing dependencies and removing unused ones,<br/>
`go mod tidy`<br/>
updates go.sum, [More Info](https://stackoverflow.com/questions/71495619/difference-between-go-mod-download-and-go-mod-tidy)<br/>

For making a vendored copy of your dependencies,<br/>
`go mod vendor`<br/>
Creates a local copy of all the dependencies so that our code works even if the dependencies are taken down from the original source, [More Info](https://stackoverflow.com/questions/68544611/what-is-the-purpose-of-go-mod-vendor-command)<br/>
