# Cron Parser
Write a command line application or script that parses a cron string and expands each field to show the times at which it will run. <br>
You should only consider the standard cron format with five time fields (minute, hour, day of month, month, and day of week) plus a command, and you do not need to handle the special time strings such as "@yearly". The input will be on a single line.<br>
The cron string will be passed to your application as a single argument.<br>
~$ your-program "d"<br>
The output should be formatted as a table with the field name taking the first 14 columns and the times as a space-separated list following it.<br>
For example, the following input argument:<br>
*/15 0 1,15 * 1-5 /usr/bin/find<br>
Should yield the following output:<br>
minute 0 15 30 45<br>
hour 0<br>
day of month 1 15<br>
month 1 2 3 4 5 6 7 8 9 10 11 12<br>
day of week 1 2 3 4 5<br>
command /usr/bin/find<br>

# To run this code: ./cron
# to run tests : Go to any of the test files, go above test function, click run tests
