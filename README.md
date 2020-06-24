# SDNS Example Plugin

### Compiling the Plugin
The plugin package is compiled using the Go toolchain. The only requirement is to use the `buildmode=plugin` compilation flag as shown below:

```
go build -buildmode=plugin -o exampleplugin.so
```