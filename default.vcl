vcl 4.0;

backend default {
    .host = "workerbee";
    .port = "8080";
}

sub vcl_recv {
    if (req.http.Authorization || req.method != "GET") {
        return (pass);
    }

    return (hash);
}

sub vcl_backend_response {
    
}

sub vcl_deliver {
    if (obj.hits > 0) {
        set resp.http.X-Cache = "HIT";
    } else {
        set resp.http.X-Cache = "MISS";
    }

    return (deliver);
}