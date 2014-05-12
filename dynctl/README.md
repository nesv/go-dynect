# dynctl

A command-line tool for working with DynECT, that uses your `~/.netrc` file for
credentials.

## ~/.netrc

By default, dynctl looks for a record in your `~/.netrc` file with the `machine`
name `api.dynect.net`; however, you can tell dynctl to use a different record by
specifying the `--machine, -m` flag.

## Usage

    NAME:
       dynctl - utility for working with DynECT's API
    
	USAGE:
	   dynctl [global options] command [command options] [arguments...]

	VERSION:
	   0.1.0

	COMMANDS:
	   list-zones, lsz	list all zones
       help, h		Shows a list of commands or help for one command

	GLOBAL OPTIONS:
		--customer, -c 			        DynECT customer name
		--machine, -m 'api.dynect.net'  netrc machine name to refer to for credentials
		--zone, -z 				        the DNS zone
		--porcelain, -p			        enable porcelain output for piping
		--version, -v			        print the version
		--help, -h				        show help

## Commands

All sub-commands require you to provide the `--customer, -c` option.

## License

The license used for this project is the Apache License, version 2.0.

Details are in LICENSE.
