vcl 4.0;

backend default {
    .host = "workerbee";
    .port = "8081";
}

sub vcl_recv {
    if (req.method != "GET" && req.method != "HEAD") {
        return (pass);
    }

    return (hash);
}

sub vcl_backend_response {
    if (beresp.status != 200) {
        set beresp.ttl = 0s;
        return (deliver);
    }

    if (bereq.http.Authorization) {
        set beresp.http.Vary = "Authorization";
    }

    if (beresp.http.Surrogate-Key) {
        set beresp.http.X-Surrogate-Key = beresp.http.Surrogate-Key;
    }

    if (beresp.http.Content-Type ~ "application/json") {
        set beresp.http.Content-Type = "application/json; charset=utf-8";
    }

    if (bereq.method == "POST" || bereq.method == "PUT" || bereq.method == "DELETE") {
        if (beresp.http.Surrogate-Key) {
            ban("obj.http.X-Surrogate-Key ~ " + beresp.http.Surrogate-Key);
        }
    }

    if (bereq.method == "GET" || bereq.method == "HEAD") {
        set beresp.ttl = 1h;
    } else {
        set beresp.ttl = 0s;


    }
}

sub vcl_deliver {
    set resp.http.Via = "login-cache";

    if (obj.hits > 0) {
        set resp.http.X-Cache = "HIT";
    } else {
        set resp.http.X-Cache = "MISS";
    }

    return (deliver);
}