# KloudDB_Shield

[![GitHub Release][release-img]][release]
[![Go Report Card][report-card-img]][report-card]
[![Go Reference](https://pkg.go.dev/badge/github.com/klouddb/klouddbshield.svg)](https://pkg.go.dev/github.com/klouddb/klouddbshield)
[![Go Build](https://github.com/klouddb/klouddbshield/actions/workflows/release.yml/badge.svg)](https://github.com/klouddb/klouddbshield/actions/workflows/release.yml)
[![Go Vuln Check](https://github.com/klouddb/klouddbshield/actions/workflows/govulncheck.yml/badge.svg)](https://github.com/klouddb/klouddbshield/actions/workflows/govulncheck.yml)

[release-img]: https://img.shields.io/github/release/klouddb/klouddbshield.svg?logo=github
[release]: https://github.com/klouddb/klouddbshield/releases
[report-card-img]: https://goreportcard.com/badge/github.com/klouddb/klouddbshield
[report-card]: https://goreportcard.com/report/github.com/klouddb/klouddbshield


## How to run this tool on my server ?

## !! Important !!  Please refer to https://klouddb.gitbook.io/klouddb_shield for detailed documentation


KloudDB Shield serves as a comprehensive security tool designed specifically for Postgres databases, conducting around 100 essential security checks. KloudDB Shield also offers the following additional features. Here is an example of an HTML report generated by our tool https://klouddb.io/klouddbshield_allchecks_aug12.html   :

	* PII Scanner
 
	* HBA Scanner

	* Transaction wraparound detector

	* Inactive Users

	* Unique IPs

	* RDS/Aurora Security reports

	* Password Generator

	* Common Username detector

	* Password attack simulator

	* Pawned password detector

	* Inactive hba lines detector

        

You can directly download the package from releases section of repo and install the package (for example - rpm for centos and deb package for Ubuntu etc..) . You also need to edit config file after installing the package(see above mentioned blog post for detailed walkthrough)


```bash
# Centos
$ rpm -i <ciscollector file>.rpm

# Debian
$ dpkg -i <ciscollector file>.deb

Usage of ciscollector:
  -r    Run
  -version
        Print version
$ ciscollector -r
Section 1  - Operating system          - 1/6  - 16.67%
Section 2  - Installation and Planning - 4/10 - 40.00%
Section 3  - File Permissions          - 2/9  - 22.22%
Section 4  - General                   - 5/7  - 71.43%
Section 6  - Auditing and Logging      - 2/3  - 66.67%
Section 7  - Authentication            - 4/6  - 66.67%
Section 8  - Network                   - 0/2  - 0.00%
Section 9  - Replication               - 0/2  - 0.00%
Overall Score - 18/45 - 40.00%
secreport.json file generated
```

## How to run locally(without installing a package) ?

Install and run locally the server

```bash
$ go build -o ./ciscollector ./cmd/ciscollector
# Edit kshieldconfig.toml at path /etc/klouddbshield/kshieldconfig.toml 
$ ./ciscollector -r
```
## RDS Checks

Make sure you have properly configured your AWS-CLI with a valid Access Key and Region or declare AWS variables properly. NOTE - You need to run this tool from bastion host or from some place where you have access to your RDS instances(It only needs basic aws rds describe priivs and sns read privs )
```
export AWS_ACCESS_KEY_ID="ASXXXXXXX"
export AWS_SECRET_ACCESS_KEY="XXXXXXXXX"
export AWS_SESSION_TOKEN="XXXXXXXXX"
export AWS_REGION="XXXXXXXXX"
```

## [Sample config file](https://github.com/klouddb/klouddbshield/blob/main/kshieldconfig_example.toml)
Below is sample file - If you are checking for postgres comment out the mysql section or if you are only checking mysql part , comment out the postgres part. Location of the config file is /etc/klouddbshield

```

[postgres]
host="localhost" 
port="5432" 
user="postgres"
dbname="postgres"
password="xxxxx" 
maxIdleConn = 2
maxOpenConn = 2

[app]
debug = true

```

## Postgres CIS Benchmark checks covered

```
	
Section 1: Installation and Patches	
1.2	Ensure systemd Service Files Are Enabled	
1.3	Ensure Data Cluster Initialized Successfully

Section 2: Directory and File Permissions	
2.1	Ensure the file permissions mask is correct

Sectioin 3: Logging Monitoring and Auditing	
3.1.2	Ensure the log destinations are set correctly	
3.1.3	Ensure the logging collector is enabled	
3.1.4	Ensure the log file destination directory is set correctly	
3.1.5	Ensure the filename pattern for log files is set correctly	
3.1.6	Ensure the log file permissions are set correctly	
3.1.7	Ensure 'log_truncate_on_rotation' is enabled	
3.1.8	Ensure the maximum log file lifetime is set correctly	
3.1.9	Ensure the maximum log file size is set correctly	
3.1.10	Ensure the correct syslog facility is selected	
3.1.11	Ensure syslog messages are not suppressed	
3.1.12	Ensure syslog messages are not lost due to size	
3.1.13	Ensure the program name for PostgreSQL syslog messages is correct	
3.1.14	Ensure the correct messages are written to the server log	
3.1.15	Ensure the correct SQL statements generating errors are recorded	
3.1.16	Ensure 'debug_print_parse' is disabled	
3.1.17	Ensure 'debug_print_rewritten' is disabled	
3.1.18	Ensure 'debug_print_plan' is disabled	
3.1.19	Ensure 'debug_pretty_print' is enabled	
3.1.20	Ensure 'log_connections' is enabled	
3.1.21	Ensure 'log_disconnections' is enabled	
3.1.22	Ensure 'log_error_verbosity' is set correctly	
3.1.23	Ensure 'log_hostname' is set correctly	
3.1.24	Ensure 'log_line_prefix' is set correctly	
3.1.25	Ensure 'log_statement' is set correctly	
3.1.26	Ensure 'log_timezone' is set correctly	
3.2	Ensure the PostgreSQL Audit Extension (pgAudit) is enabled

Section 4: User Access and Authorization	
4.2	Ensure excessive administrative privileges are revoked	
4.3	Ensure excessive function privileges are revoked	
4.4	Ensure excessive DML privileges are revoked	
4.5	Ensure Row Level Security (RLS) is configured correctly	
4.6	Ensure the set_user extension is installed	
4.7	Make use of predefined roles

Section 5: Connection and Login	
5.1	Ensure login via "local" UNIX Domain Socket is configured correctly	
5.2	Ensure login via "host" TCP/IP Socket is configured correctly	
5.3	Ensure Password Complexity is configured

Section 6: Postgres Settings	
6.2	Ensure 'backend' runtime parameters are configured correctly	
6.3	Ensure 'Postmaster' Runtime Parameters are Configured	
6.4	Ensure 'SIGHUP' Runtime Parameters are Configured	
6.5	Ensure 'Superuser' Runtime Parameters are Configured	
6.6	Ensure 'User' Runtime Parameters are Configured	
6.7	Ensure FIPS 140-2 OpenSSL Cryptography Is Used	
6.8	Ensure SSL is enabled and configured correctly	
6.9	Ensure the pgcrypto extension is installed and configured correctly

Section 7: Replication	
7.1	Ensure a replication-only user is created and used for streaming replication	
7.2	Ensure logging of replication commands is configured	
7.3	Ensure base backups are configured and functional	
7.4	Ensure WAL archiving is configured and functional	
7.5	Ensure streaming replication parameters are configured correctly	

Section 8: Special Configuration Considerations	
8.1	Ensure PostgreSQL subdirectory locations are outside the data cluster	
8.2	Ensure the backup and restore tool, 'pgBackRest', is installed and configured	
8.3	Ensure miscellaneous configuration settings are correct
```
## Contributing 

We welcome PRs and Issue Reports

## Help 

Please reach us at support@klouddb.io 

