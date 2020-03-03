# goshowips
Pipe in a file or program output and output all unique IP addresses, one per line

## Example
[luke@somehost]$ cat ~/scans/nmap/XXX/nmapoutput.gnmap | grep open | ./goshowips  
192.168.1.1  
192.168.1.2  
192.168.1.3  
192.168.1.4  
192.168.1.5  
192.168.1.6  
192.168.1.7  

## Example with GeoIP Information
[luke@somehost]$ cat ~/scans/nmap/XXX/nmapoutput.gnmap | grep open | ./goshowips -g
192.168.1.1  US  
192.168.1.2  CN  
192.168.1.3  GB  
192.168.1.4  IT  
192.168.1.5  IR  
192.168.1.6  US  
192.168.1.7  CA  

Obviously these IP addresses would not resolve to the countries listed. :)
