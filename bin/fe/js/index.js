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
		return colection;
	}
});

app.Phone_View = Backbone.View.extend({
	template : _.template($('#phone-tview').html()),
	initialize : function() {
		this.setElement(this.template(this.model));
	},
	render : function() {
		var elem = this.el.lastElementChild
		if (elem.innerText == "valid"){
			$(elem).addClass("valid");
		}else{
			$(elem).addClass("nvalid");
		}
		return this;
	}
});

app.table = Backbone.View.extend({
	initialize : function() {
		this.el = document.getElementsByClassName("table-content")[0];
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
		this.phones = new app.Phone_Adds().fetch();
		return this;
	},
	events : {
		'keyup #myInput' : 'filterCountry',
		'click #state' : 'cycleState'
	},
	filterCountry : function(e){
		var elems = $(".table-row").hide();
		var filtered_items = _.filter(elems,function(elem){
			return elem.innerText.indexOf(e.target.value) > -1;
		});
		$(filtered_items).show();
	},
	cycleState : function(e){
		var valid = $(".valid");
		var nvalid = $(".nvalid");
		if (valid[0].parentElement.style.display == "" && nvalid[0].parentElement.style.display == ""){
			for (cell of nvalid){
				$(cell.parentElement).hide();
			}
		}else if(valid[0].parentElement.style.display == "" && nvalid[0].parentElement.style.display == "none"){
			for (cell of nvalid){
				$(cell.parentElement).show();
			}
			for (cell of valid){
				$(cell.parentElement).hide();			
			}
		}else{
			for (cell of nvalid){
				$(cell.parentElement).show();
			}
			for (cell of valid){
				$(cell.parentElement).show();
			}
		}
	}
	
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
