
# Função

O fsortd serve para organizar os arquivos e evitar a bagunça que a pasta de dowloand fica




# Planejamento

Monitoramento: fsnotify = fsnotify is a Go library to provide cross-platform filesystem notifications on Windows, Linux, macOS, BSD, and illumos.
regras: TOML
logs: slog
cli: cobra





# TOML estructure
Delay = minutes
dest = Destination

```TOML
[Config]
Source = "~/Downloads"
Delay = 4 
AutoMakeDir = false



[Image]
FileFormat = ["png","jpg"]
dest = "~/Pictures"
```


