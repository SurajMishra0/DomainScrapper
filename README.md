# DomainScrapper
This script is a domain grabber tool that allows users to input a TLD and retrieve the available domains for that TLD by scraping the "https://zoxh.com/tld" website.
The retrieved domain names are stored in separate files based on their TLDs.

# Necessary modules and libraries For DomainScrapper

requests,
re,
pyfiglet,
bs4,
os,
datetime,

# Working

1. The Python script will ask you the TLD which you want to scrap from internet, After entering desired TLD Script checks if a given top-level domain (TLD) exists by scraping the "https://zoxh.com/tld" webpage. It retrieves the list of available TLDs and checks if the given domain parameter is present in that list. It returns True if the TLD exists and False otherwise. If the TLD is not valid, it displays an error message and exits the script.
2. Then it performs the actual domain grabbing process. It takes a TLD as input and scrapes the "https://zoxh.com/tld/{domain_tld}" webpage to retrieve the total number of domains available for that TLD. It then iterates through each page of the TLD's domains and grabs the domain names. The domain names are stored in a file named "tld_{domain_tld}.txt". The progress is printed on the console, indicating the number of domains grabbed and the current page.

[It search thorugh approx 200 pages for desirect tld.]

# We are using GO lang to convert domains into ip.

# Requirements 
1. Install Go: If you haven't already, download and install the Go programming language from the official Go website: https://golang.org/dl/
2. Check if Go is installed succesfully by running " go version "

# Run Script

go build Domain2IP.go (not needed actually direct run is enough)
go run Domain2IP.go

# Range IP
 
We are ranging IP so that we can reverse that ip later and find all the domain being hosted on those ip. it will increase number of domains drastically. 

1. We already have go installed in our system so a code " go run RangeIP.go " is enough.
2. It will start ranging ip from a given list and will save in a new list.
