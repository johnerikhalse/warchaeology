## warc convert

Convert web archives to warc files. Use subcommands for the supported formats

### Options

```
  -h, --help   help for convert
```

### Options inherited from parent commands

```
      --config string          config file. If not set, /etc/warc/, $HOME/.warc/ and current working dir will be searched for file config.yaml
      --log-console strings    The kind of log output to write to console. Valid values: info, error, summary, progress (default [progress,summary])
      --log-file strings       The kind of log output to write to file. Valid values: info, error, summary (default [info,error,summary])
  -L, --log-file-name string   a file to write log output. Empty for no log file
      --tmpdir string          directory to use for temporary files (default "/tmp")
```

### SEE ALSO

* [warc](warc.md)	 - A tool for handling warc files
* [warc convert arc](warc_convert_arc.md)	 - Convert arc file into warc file
* [warc convert nedlib](warc_convert_nedlib.md)	 - Convert directory with files harvested with Nedlib into warc files

