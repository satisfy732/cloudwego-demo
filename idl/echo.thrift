namespace go api

struct Request {
    string message
}

struct Reponse {
    string message
}

service Echo {
    Reponse echo(Request req)
}
