# localhost-proxy

To dev and run with WebRTC/MediaStream - you need to run on localhost or with a self-signed certificate. To work around SSL cert, I've created a simple proxy that will run on localhost and proxy all requests to the specified host.

Usage of proxy:
  -l, --listen string   listening IP (default "0.0.0.0")
  -p, --port int        listen on port (default 80)
  -t, --target string   target URL

Example:
  proxy -t http://remoteIP:3000
  
Then use your morden browser to access localhost:80 and you will be redirected to the target URL, and you will find you can acquire a MediaStream (Mic or Cam) from the target URL.

Enjoy this handy little tool!
