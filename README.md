# Nessus_Charts
Nessus is an excellent vulnerability scanner but a HORRIBLE reporting tool.

Hopefully this can make your life a *little* bit easier.

## Usage Instructions
1. Download the binary for your Operating System [Linux](https://github.com/tbcsec/Nessus_Charts/releases/download/v0.1-alpha/nessus_charts) or [Windows](https://github.com/tbcsec/Nessus_Charts/releases/download/v0.1-alpha/nessus_charts.exe)
2. Download the [Excel template](https://github.com/tbcsec/Nessus_Charts/releases/download/v0.1-alpha/Chart_Template.xlsx)
3. Run the binary and provide the filepaths to the CSV you would like to import `-csv=my_scan.csv`, the Excel template file `-excel=Chart_Template.csv`, the SQLite3 database file `-sql=new_db.db` (we will create one at the provided path if a db file does not exist yet), and the table name you would like to use `-table="Internal Scan Jan 01"`.

For a list of the variables, simply run the binary with `-h`. This will also show you the default options. None of the options are required, so for the simplest use, just put the binary, CSV file, and Excel template in the same directory. Your command could look something like `./nessus_charts -csv=nessus_export.csv`

Relative file paths or full file paths are supported. If any of your options have spaces in them, be sure to enclose your option in "quotes".

Here is a complete sample run for a Windows system:
```
nessus_charts.exe -csv="c:\Users\me\Downloads\Really big scan.csv" -sql=..\Documents\sqlite3_rocks.db -excel=Chart_Template.xlsx -table="External AWS Environment"
```

One last note, make sure to make your table names unique each time you run the program to avoid duplicate entries in a single table which could provide falsely high vulnerability data.

## Why Does This Exist?
TBConsulting does lots of security engagements which involve vulnerability detection and reporting. These are one-off engagements, not vulnerability management. We wanted to provide customers with meaningful insights about the data collected in Nessus, but the only options in Nessus are CSV export and HTML.

The HTML export will provide you with a webpage that has no end. The CSV export also provides just the raw data and leaves you to interpret. We started off like I am sure many have, with Excel and pivot table fun. The problem we had with that is some of the customers we are scanning have thousands of hosts and hundreds of thousands of vulnerabilities. I hope nobody is offended, but Excel is not actually a database...

So, after spending hours with each CSV file from each scan we did in Nessus and watching our machines run out of memory and crash while Excel tried to read a 5GB CSV file; we decided a solution was needed. Tenable clearly wants you to purchase either Tenable.io or Tenable.sc. Those are both great products, and we use them for ongoing vulnerability management, but they are not meant for one-off scans. You do not want an asset that you scanned one time to count against your ongoing asset count for licensing.

This simple tool solves this problem by reading your provided CSV, writing all those records to a SQLite database, running 5 SQL queries, and then writing the results to our template Excel doc which fills out some pre-built charts.

If you have any questions, you can email us [here](mailto:security@tbconsulting.com)

We would love to hear what else we can do to make this tool more helpful to you!