# logr-survey

This is a simple tool to explore how different loggers have different logr V level defaults


| logger\nimplementation | Default Output | Max Info Level | Max Error Level |
| ---------------------- | -------------- | -------------- | --------------- |
| logr.Discard()         |          None  |          None  |            None |
| glogr                  |        Stderr  |          None  |               5 |
| klogr                  |        Stderr  |             0  |               5 |
| zapr (Production)      |        Stderr  |             0  |               5 |
| stdr                   |      User Set  |             0  |               5 |
| logrusr                |        Stderr  |             0  |               5 |
| logfmtr                |        Stdout  |             0  |               5 |
| zerologr               |      User Set  |             2  |               5 |