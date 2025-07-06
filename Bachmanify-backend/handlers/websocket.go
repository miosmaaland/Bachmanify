package handlers

import (
    "log"
    "net/http"
    "strings"

    "github.com/gorilla/websocket"
)

type Node struct {
    IP         string
    Vulnerable bool
    Exploited  bool
    Connected  bool
}

var network = []Node{
    {"192.168.1.10", true, false, false},
    {"192.168.1.20", false, false, false},
    {"192.168.1.30", true, false, false},
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket upgrade failed:", err)
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
        input := strings.TrimSpace(string(msg))
        response := handleCommand(input)
        err = conn.WriteMessage(websocket.TextMessage, []byte(response))
        if err != nil {
            log.Println("Write error:", err)
            break
        }
    }
}

func handleCommand(cmd string) string {
    log.Println("Received command:", cmd)
    parts := strings.Fields(strings.ToLower(strings.TrimSpace(cmd)))
    if len(parts) == 0 {
        return "Please enter a command."
    }

    switch parts[0] {
    case "help":
        return "Available commands: scan, connect <IP>, exploit <IP>"

    case "scan":
        found := []string{}
        for _, node := range network {
            if node.Vulnerable {
                found = append(found, node.IP)
            }
        }
        if len(found) == 0 {
            return "No vulnerable nodes found."
        }
        return "Vulnerable nodes found: " + strings.Join(found, ", ")

    case "connect":
        if len(parts) < 2 {
            return "Usage: connect <IP>"
        }
        ip := parts[1]
        for i, node := range network {
            if node.IP == ip {
                if node.Connected {
                    return "Already connected to " + ip
                }
                network[i].Connected = true
                return "Connected to " + ip
            }
        }
        return "Node " + ip + " not found."

    case "exploit":
        if len(parts) < 2 {
            return "Usage: exploit <IP>"
        }
        ip := parts[1]
        for i, node := range network {
            if node.IP == ip {
                if !node.Connected {
                    return "Not connected to " + ip + ". Connect first."
                }
                if node.Vulnerable && !node.Exploited {
                    network[i].Exploited = true
                    return "Exploit successful on " + ip
                }
                if node.Exploited {
                    return ip + " is already exploited."
                }
                return "Exploit failed: node not vulnerable."
            }
        }
        return "Node " + ip + " not found."

    default:
        return "Unknown command. Type 'help' for available commands."
    }
}

