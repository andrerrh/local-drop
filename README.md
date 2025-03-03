# local-drop  
A local network file server built with Go and HTMX.

## Overview  
**local-drop** is a lightweight file server that allows file sharing over a local network.  
It is built using:  
- **Go** with Go HTML templates for backend processing  
- **HTMX** and a bit of **JavaScript** for frontend interactivity  
- **Tailwind CSS (CDN)** for styling  

The project supports **drag-and-drop** file uploads, including multiple files at once.

## Installation & Usage  
1. **Clone the repository**  
   git clone https://github.com/yourusername/local-drop.git  
   cd local-drop  

2. **Run the server**  
   go run main.go  

3. **Access the web interface**  
   - Open http://localhost:8080 on your browser  
   - Or use http://server-ip:8080 from another device on the same network  

## Configuration  
A custom port can be set in the `config.json` file (defaults to `8080`).  
{
  "port": 8080
}
