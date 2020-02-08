var app = {}; // create namespace for our app

app.Phone_Add = Backbone.Model.extend({
	defaults : {
		id : 0,
		name : "",
		number : 0,
		country : "",
		state : ""
	}
});

app.Phone_Adds = Backbone.Collection.extend({
	model : app.Phone_Add,
	url : '/customer',
	parse : function(colection) {
		this.phoneview = new app.table({
			model : colection
		}).render();
	}
});

app.Phone_View = Backbone.View.extend({
	template : _.template($('#phone-tview').html()),
	initialize : function() {
		this.setElement(this.template(this.model));
	},
	render : function() {
		elem = this.el.lastElementChild
		if (elem.innerText == "valid"){
			elem.className = "valid";
		}else{
			elem.className = "nvalid";
		}
		return this;
	}
});

app.table = Backbone.View.extend({
	el : '#thetable',
	initialize : function() {

	},
	render : function() {
		var self = this;
		this.model.forEach(function(val, i) {
			self.el.append(new app.Phone_View({
				model : {
					name : val.name,
					number : val.number,
					country : val.country,
					state : (val.state ? "valid" : "not valid")
				}
			}).render().el);
		});
	}
});

// main view
app.main = Backbone.View.extend({
	el : '#content',
	initialize : function() {
		this.phones = new app.Phone_Adds();
		this.phones.fetch();
		return this;
	},
});

app.Router = Backbone.Router.extend({
	routes : {

	}
});

// --------------
// Initializers
// --------------
Backbone.history.start();
var content = new app.main();

function filterCountry() {
	// Declare variables
	var input, filter, table, tr, td, i, txtValue;
	input = document.getElementById("myInput");
	filter = input.value.toUpperCase();
	table = document.getElementById("thetable");
	tr = table.getElementsByTagName("tr");

	// Loop through all table rows, and hide those who don't match the search
	// query
	for (i = 0; i < tr.length; i++) {
		td = tr[i].getElementsByTagName("td")[2];
		if (td) {
			txtValue = td.textContent || td.innerText;
			if (txtValue.toUpperCase().indexOf(filter) > -1) {
				tr[i].style.display = "";
			} else {
				tr[i].style.display = "none";
			}
		}
	}
}

function cycleState(){
	valid = document.getElementsByClassName("valid");
	nvalid = document.getElementsByClassName("nvalid");
	if (!valid[0].parentElement.hidden && !nvalid[0].parentElement.hidden){
		for (cell of nvalid){
			cell.parentElement.hidden = true;
		}
	}else if(!valid[0].parentElement.hidden && nvalid[0].parentElement.hidden){
		for (cell of nvalid){
			cell.parentElement.hidden = false;
		}
		for (cell of valid){
			cell.parentElement.hidden = true;
		}
	}else{
		for (cell of nvalid){
			cell.parentElement.hidden = false;
		}
		for (cell of valid){
			cell.parentElement.hidden = false;
		}
	}
}
