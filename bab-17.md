# GopherJS

##membuat package main
~~~bash

package main

import (
  "github.com/gopherjs/gopherjs/js"
  "github.com/rolaveric/pet"
)

func main() {
	js.Global.Set("pet", map[string]interface{}{
		"New": pet.New,
	})
}
~~~

##membuat package pet
~~~bash
package pet

import "github.com/gopherjs/gopherjs/js"

type Pet struct {
	name string
}

func (p *Pet) Name() string {
	return p.name
}

func (p *Pet) SetName(newName string) {
	p.name = newName
}

func New(name string) *js.Object {
	return js.MakeWrapper(&Pet{name})
}
~~~

##gropherjs result short.js
~~~bash
// Lines 8-13
var go$global;
if (typeof window !== "undefined") {
	go$global = window;
} else if (typeof GLOBAL !== "undefined") {
	go$global = GLOBAL;
}

// Lines 1425-1461
go$packages["github.com/rolaveric/gopherjs/pet"] = (function() {
	var go$pkg = {}, Pet, New;
	Pet = go$pkg.Pet = go$newType(0, "Struct", "pet.Pet", "Pet", "github.com/rolaveric/gopherjs/pet", function(name_) {
		this.go$val = this;
		this.name = name_ !== undefined ? name_ : "";
	});
	Pet.Ptr.prototype.Name = function() {
		var p;
		p = this;
		return p.name;
	};
	Pet.prototype.Name = function() { return this.go$val.Name(); };
	Pet.Ptr.prototype.SetName = function(newName) {
		var p;
		p = this;
		p.name = newName;
	};
	Pet.prototype.SetName = function(newName) { return this.go$val.SetName(newName); };
	New = go$pkg.New = function(name) {
		return new Pet.Ptr(name);
	};
	go$pkg.init = function() {
		(go$ptrType(Pet)).methods = [["Name", "", [], [Go$String], false, -1], ["SetName", "", [Go$String], [], false, -1]];
		Pet.init([["name", "name", "github.com/rolaveric/gopherjs/pet", Go$String, ""]]);
	}
	return go$pkg;
})();
go$packages["C:\GoCode\src\github.com\rolaveric\gopherjs"] = (function() {
	var go$pkg = {}, js = go$packages["github.com/gopherjs/gopherjs/js"], pet = go$packages["github.com/rolaveric/gopherjs/pet"], main;
	main = go$pkg.main = function() {
		var _map, _key;
		go$global.pet = go$externalize((_map = new Go$Map(), _key = "New", _map[_key] = { k: _key, v: new (go$funcType([Go$String], [(go$ptrType(pet.Pet))], false))(pet.New) }, _map), (go$mapType(Go$String, go$emptyInterface)));
	};
	go$pkg.init = function() {
	}
	return go$pkg;
})();
~~~

##Posting perpustakaan Javascript untuk Go
Katakanlah kita punya ' Pengguna objek model yang menggunakan variabel global ' DB ' untuk membuat panggilan database SQL 
###gropherjs model original
~~~bash
// User Type
function User(name, id) {
  this.name = name;
  this.id = id;
  
  this.save = function () {
    DB.query('UPDATE User SET name = ? WHERE id = ?', this.name, this.id);
  }
}

// Factory for creating a new user
User.new = function (name) {
  DB.query('INSERT INTO User (name) VALUES (?)', name);
  var id = DB.query('SELECT @@IDENTITY').nextRow()[0];
  return new User(name, id);
};

// Retrieves a user from the database
User.get = function (id) {
  var result = DB.query('SELECT name FROM User WHERE id = ?', id);
  if (result.rowCount === 0) {
    return null;
  }
  var name = result.nextRow()[0];
  return new User(name, id);
};

// Retrieves all users from the database
User.all = function () {
  var users = [];
  var result = DB.query('SELECT name, id FROM User');
  for (var x = 0; x < result.rowCount; x++) {
    var row = result.nextRow();
    users.push(new User(row[0], row[1]));
  }
  return users;
};

// Use case that still needs to use this model
function myHttpHandler(request) {
  if (request.method === 'GET') {
    var id = request.params['userid'];
    if (id) {
      return User.get(id);
    } else {
      return User.all();
    }
  } else if (request.method === 'POST') {
    return User.new(request.params['name']);
  }
}
~~~

##Beberapa hal yang kita tahu akan berbeda ketika kita mengubahnya ke Go .
API refactored terlihat seperti :

~~~bash
// Namespace created within an IIFE for private scope
var user = (function () {
  // Variable for holding the injected DB interface
  var DB;
  
  // User Type
  function User(name, id) {/* ... */}
  
  return {
    // Expose a function for setting the DB interface
    "registerDB": function (db) {
      DB = db;
    },
    "new": function (name) {
      // The "DB" type's methods will be capitalised
      DB.Query('INSERT INTO User (name) VALUES (?)', name);
      
      // Lets be a bit more type safe and specify that we expect an int
      var id = DB.Query('SELECT @@IDENTITY').NextRow().GetInt(0);
      
      return new User(name, id);
    },
    "get": function (id) {/* ... */},
    "all": function () {/* ... */},
    "save": function () {/* ... */}
  };
})();
~~~

##kode go
~~~
package user

// Interface for a database result row
type DBRow interface {
	GetInt(colnum int) int
	GetString(colnum int) string
}

// Interface for a database result
type DBResult interface {
	NextRow() DBRow
	RowCount() int
}

// Interface for an object which can be used to make database queries
type DB interface {
	Query(query string, params ...interface{}) DBResult
}

// Private package variable for the registered DB interface
var db DB

// Method for registering a DB interface
func RegisterDB(newDb DB) {
	db = newDb
}

// User type
type User struct {
	Name string
	ID   int
}

// Save method for the User type
func Save(u *User) {
	db.Query("UPDATE User SET name = ? WHERE id = ?", u.Name, u.ID)
}

// Function for creating a new User
func New(name string) *User {
	db.Query("INSERT INTO User (name) VALUES (?)", name)
	id := db.Query("SELECT @@IDENTITY").NextRow().GetInt(0)
	return &User{name, id}
}

// Function for getting a single User
func Get(id int) *User {
	result := db.Query("SELECT name FROM User WHERE id = ?", id)
	if result.RowCount() == 0 {
		return nil
	}
	name := result.NextRow().GetString(0)
	return &User{name, id}
}

// Function for getting all users
func All() []*User {
	result := db.Query("SELECT name, id FROM User")
	users := make([]*User, result.RowCount())
	for x, c := 0, result.RowCount(); x < c; x++ {
		row := result.NextRow()
		users[x] = &User{row.GetString(0), row.GetInt(1)}
	}
	return users
}
~~~

##model main go

~~~
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rolaveric/gopherjs-demo/user"
	"github.com/rolaveric/gopherjs-demo/user/js/db"
)

// Starting point for compiling JS code
func main() {
	js.Global.Set("user", map[string]interface{}{
		"registerDB": RegisterDBJS,
		"new":        user.New,
		"get":        user.Get,
		"all":        user.All,
		"save":       SaveJS,
	})
}

// Takes a DB adapter written in Javascript and wraps it as a DB interface
func RegisterDBJS(o *js.Object) {
	user.RegisterDB(db.JSDB{o})
}

// Takes a JS object and wraps it as a User struct
func SaveJS(o *js.Object) {
	user.Save(&user.User{o.Get("Name").String(), o.Get("ID").Int()})
}
~~~