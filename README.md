# Heffalump
Heffalump is an endless honeypot that gives malicious bots nightmares. To use, in your robots.txt tell robots not to go to a certain URL, which heffalump is reverse proxying. Any web agent that does go to the URL will receive an endless stream of random data, which will overflow its memory and/or storage if it doesn't have a max buffer size set or at the very least severely waste its time.

The source of the honeypot data is [Once On a Time](http://www.gutenberg.org/files/27771/27771-h/27771-h.htm), one of A. A. Milne's most beloved and most public domain works.

## Usage
Usage of heffalump:

heffalump [<network address> [<path>]]

    heffalump serves an endless HTTP honeypot

    <network address> defaults to ":8080".

    <path> defaults to "/". Paths ending in "/" will match all sub-pathes.
