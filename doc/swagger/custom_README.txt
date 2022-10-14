TODO list

1. update swagger-initializer.js
    change the value of "url" to the swagger json file

2. edit router
	fs := http.FileServer(http.Dir("./doc/swagger"))
	mux.Handle(
        "/swagger/",
        http.StripPrefix("/swagger/", fs), // "/swagger/" must be stripped to navigate to the correct dir.
    )

3. access localhost/swagger
