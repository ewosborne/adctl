# adctl

`adctl` controls AdGuard Home.


## Usage
    adctl [command]

    Available Commands:
    completion  Generate the autocompletion script for the specified shell
    disable     Disable ad blocker. Optional duration in time.Duration format.
    enable      Enable ad blocking
    getlog      Get logs. Optional length parameter, 0 == MaxUint32 log length..
    help        Help about any command
    status      Get adblock status
    toggle      Toggle adblocker between enabled and disabled.

You need three environment variables: 

    ADCTL_USERNAME
    ADCTL_PASSWORD
    ADCTL_HOST

The username and password are what you'd use to log into the AdGuard Home console. `ADCTL_HOST` is the host and port you use to reach the GUI.  Mine is set to `router:8080` but IP address will work too. AdGuard Home doesn't support auth tokens so hardcoded password is all you get. 

## Building
I use [just](https://just.systems/) so everything is in a `justfile`. You can do that too, or just copy the steps from that file and run them by hand. No magic here. `just build` or `go build -ldflags "-s -w" .`, as you like.


## Examples

### enable
    eric@air ~ % adctl enable
    adguard is enabled

### disable [timestamp]
Disables adblocker. Takes an optional [duration string](https://pkg.go.dev/time#ParseDuration).

    eric@air ~ % adctl disable
    adguard is disabled
    eric@air ~ % adctl disable 30s
    adguard is disabled for 29s

### getlog
Returns a json of the logs from the server. The server's default is 500 entries (0-499) but you can limit this with `getlog [limit]`.  

    eric@air ~ % adctl getlog 42 | jq '.data[41]'
    {
    "answer": [
        {
        "type": "A",
        "value": "0.0.0.0",
        "ttl": 10
        }
    ],
    "answer_dnssec": false,
    "cached": false,
    "client": "192.168.1.182",
    "client_info": {
        "whois": {},
        "name": "",
        "disallowed_rule": "192.168.1.182",
        "disallowed": false
    },
    "client_proto": "",
    "elapsedMs": "0.13885699999999998",
    "filterId": 1732762630,
    "question": {
        "class": "IN",
        "name": "mobile.events.data.microsoft.com",
        "type": "A"
    },
    "reason": "FilteredBlackList",
    "rule": "||events.data.microsoft.com^",
    "rules": [
        {
        "filter_list_id": 1732762630,
        "text": "||events.data.microsoft.com^"
        }
    ],
    "status": "NOERROR",
    "time": "2024-12-08T05:55:33.752856831-05:00",
    "upstream": ""
    }


A limit of 0 returns all logs on the server.

    eric@air adctl % adctl getlog 0 | gron | tail   
    json.data[50394].question = {};
    json.data[50394].question.name = "lb._dns-sd._udp.example.com";
    json.data[50394].question.type = "PTR";
    json.data[50394].question["class"] = "IN";
    json.data[50394].reason = "NotFilteredNotFound";
    json.data[50394].rules = [];
    json.data[50394].status = "NXDOMAIN";
    json.data[50394].time = "2024-12-07T13:34:17.590528405-05:00";
    json.data[50394].upstream = "127.0.0.1:5353";
    json.oldest = "2024-12-07T13:34:17.590528405-05:00";


### status

    eric@air ~ % adctl status
    adguard is enabled

### toggle
Flips from whatever status it is now (enabled | disabled) to the other one. 

    eric@air adctl % adctl status
    adguard is enabled
    eric@air adctl % adctl toggle
    adguard is disabled
    eric@air adctl % adctl status
    adguard is disabled
    eric@air adctl % adctl toggle
    adguard is enabled
    eric@air adctl % adctl status
    adguard is enabled

If a duration is set on a disable then toggling does not preserve it.

    eric@air adctl % adctl disable 1m
    adguard is disabled for 59s
    eric@air adctl % adctl toggle
    adguard is enabled
    eric@air adctl % adctl toggle
    adguard is disabled
    eric@air adctl % adctl status
    adguard is disabled