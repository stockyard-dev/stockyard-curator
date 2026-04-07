package server

import "net/http"

func (s *Server) dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(dashHTML))
}

const dashHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1.0">
<title>Curator</title>
<link href="https://fonts.googleapis.com/css2?family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&family=JetBrains+Mono:wght@400;500;700&display=swap" rel="stylesheet">
<style>
:root{--bg:#1a1410;--bg2:#241e18;--bg3:#2e261e;--rust:#e8753a;--leather:#a0845c;--cream:#f0e6d3;--cd:#bfb5a3;--cm:#7a7060;--gold:#d4a843;--green:#4a9e5c;--red:#c94444;--mono:'JetBrains Mono',monospace;--serif:'Libre Baskerville',serif}
*{margin:0;padding:0;box-sizing:border-box}
body{background:var(--bg);color:var(--cream);font-family:var(--serif);line-height:1.6}
.hdr{padding:.9rem 1.5rem;border-bottom:1px solid var(--bg3);display:flex;justify-content:space-between;align-items:center;gap:1rem;flex-wrap:wrap}
.hdr h1{font-family:var(--mono);font-size:.9rem;letter-spacing:2px}
.hdr h1 span{color:var(--rust)}
.main{padding:1.2rem 1.5rem;max-width:1100px;margin:0 auto}
.stats{display:grid;grid-template-columns:repeat(4,1fr);gap:.5rem;margin-bottom:1.2rem;font-family:var(--mono)}
.st{background:var(--bg2);border:1px solid var(--bg3);padding:.7rem;text-align:center}
.st-v{font-size:1.3rem;font-weight:700;color:var(--gold)}
.st-l{font-size:.5rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-top:.2rem}
.toolbar{display:flex;gap:.5rem;margin-bottom:1rem;flex-wrap:wrap;align-items:center}
.search{flex:1;min-width:180px;padding:.4rem .6rem;background:var(--bg2);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}
.search:focus{outline:none;border-color:var(--leather)}
.filter-sel{padding:.4rem .5rem;background:var(--bg2);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.65rem}
.grid{display:grid;grid-template-columns:repeat(auto-fill,minmax(280px,1fr));gap:.8rem}
.card{background:var(--bg2);border:1px solid var(--bg3);padding:1rem;cursor:pointer;transition:border-color .15s;display:flex;flex-direction:column;gap:.4rem}
.card:hover{border-color:var(--leather)}
.card-title{font-family:var(--serif);font-size:1.05rem;font-weight:700;color:var(--cream);line-height:1.3}
.card-cat{font-family:var(--mono);font-size:.55rem;color:var(--leather);text-transform:uppercase;letter-spacing:1px}
.stars{font-family:var(--mono);font-size:.75rem;color:var(--gold);letter-spacing:1px}
.stars .empty-star{color:var(--bg3)}
.card-meta{font-family:var(--mono);font-size:.55rem;color:var(--cm);display:flex;gap:.6rem;flex-wrap:wrap}
.card-meta span{display:inline-flex;align-items:center;gap:.2rem}
.card-tags{display:flex;gap:.3rem;flex-wrap:wrap;margin-top:.2rem}
.tag{font-family:var(--mono);font-size:.5rem;background:var(--bg3);color:var(--cd);padding:.1rem .35rem}
.card-extra{font-family:var(--mono);font-size:.55rem;color:var(--cd);margin-top:.4rem;padding-top:.3rem;border-top:1px dashed var(--bg3);display:flex;flex-direction:column;gap:.15rem}
.card-extra-row{display:flex;gap:.4rem}
.card-extra-label{color:var(--cm);text-transform:uppercase;letter-spacing:.5px;min-width:80px}
.card-extra-val{color:var(--cream)}
.btn{font-family:var(--mono);font-size:.6rem;padding:.3rem .6rem;cursor:pointer;border:1px solid var(--bg3);background:var(--bg);color:var(--cd);transition:.15s}
.btn:hover{border-color:var(--leather);color:var(--cream)}
.btn-p{background:var(--rust);border-color:var(--rust);color:#fff}
.btn-p:hover{opacity:.85;color:#fff}
.modal-bg{display:none;position:fixed;inset:0;background:rgba(0,0,0,.65);z-index:100;align-items:center;justify-content:center}
.modal-bg.open{display:flex}
.modal{background:var(--bg2);border:1px solid var(--bg3);padding:1.5rem;width:560px;max-width:92vw;max-height:90vh;overflow-y:auto}
.modal h2{font-family:var(--serif);font-size:1.05rem;margin-bottom:1rem;color:var(--rust)}
.fr{margin-bottom:.6rem}
.fr label{display:block;font-family:var(--mono);font-size:.55rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-bottom:.2rem}
.fr input,.fr select,.fr textarea{width:100%;padding:.4rem .5rem;background:var(--bg);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}
.fr input:focus,.fr select:focus,.fr textarea:focus{outline:none;border-color:var(--leather)}
.fr textarea{font-family:var(--serif);line-height:1.5}
.row3{display:grid;grid-template-columns:1fr 1fr 1fr;gap:.5rem}
.row2{display:grid;grid-template-columns:1fr 1fr;gap:.5rem}
.fr-section{margin-top:1rem;padding-top:.8rem;border-top:1px solid var(--bg3)}
.fr-section-label{font-family:var(--mono);font-size:.55rem;color:var(--rust);text-transform:uppercase;letter-spacing:1px;margin-bottom:.5rem}
.rating-input{display:flex;gap:.4rem;align-items:center}
.star-btn{font-size:1.4rem;cursor:pointer;color:var(--bg3);background:none;border:none;padding:0 .1rem;line-height:1}
.star-btn.lit{color:var(--gold)}
.star-clear{font-family:var(--mono);font-size:.55rem;color:var(--cm);cursor:pointer;margin-left:.5rem}
.star-clear:hover{color:var(--cream)}
.acts{display:flex;gap:.4rem;justify-content:flex-end;margin-top:1rem}
.acts .btn-del{margin-right:auto;color:var(--red);border-color:#3a1a1a}
.acts .btn-del:hover{border-color:var(--red);color:var(--red)}
.empty{text-align:center;padding:3rem;color:var(--cm);font-style:italic;font-size:.85rem}
@media(max-width:600px){.stats{grid-template-columns:repeat(2,1fr)}.row3{grid-template-columns:1fr 1fr}}
</style>
</head>
<body>

<div class="hdr">
<h1 id="dash-title"><span>&#9670;</span> CURATOR</h1>
<button class="btn btn-p" onclick="openNew()">+ New Recipe</button>
</div>

<div class="main">
<div class="stats" id="stats"></div>
<div class="toolbar">
<input class="search" id="search" placeholder="Search title, ingredients, tags..." oninput="debouncedRender()">
<select class="filter-sel" id="category-filter" onchange="render()">
<option value="">All Categories</option>
</select>
<select class="filter-sel" id="rating-filter" onchange="render()">
<option value="0">Any Rating</option>
<option value="3">3+ stars</option>
<option value="4">4+ stars</option>
<option value="5">5 stars only</option>
</select>
</div>
<div id="grid" class="grid"></div>
</div>

<div class="modal-bg" id="mbg" onclick="if(event.target===this)closeModal()">
<div class="modal" id="mdl"></div>
</div>

<script>
var A='/api';
var RESOURCE='recipes';

var fields=[
{name:'title',label:'Title',type:'text',required:true},
{name:'category',label:'Category',type:'select_or_text',options:[]},
{name:'ingredients',label:'Ingredients',type:'textarea',placeholder:'One per line'},
{name:'instructions',label:'Instructions',type:'textarea'},
{name:'prep_time',label:'Prep (min)',type:'number'},
{name:'cook_time',label:'Cook (min)',type:'number'},
{name:'servings',label:'Servings',type:'number'},
{name:'tags',label:'Tags',type:'text',placeholder:'comma separated'},
{name:'rating',label:'Rating',type:'rating'}
];

var recipes=[],recipeExtras={},editId=null,searchTimer=null,currentRating=0;

// ─── Helpers ──────────────────────────────────────────────────────

function fieldByName(n){
for(var i=0;i<fields.length;i++)if(fields[i].name===n)return fields[i];
return null;
}

function debouncedRender(){
clearTimeout(searchTimer);
searchTimer=setTimeout(render,200);
}

function starsHTML(rating){
var n=Math.round(rating||0);
var h='<span class="stars">';
for(var i=1;i<=5;i++){
if(i<=n)h+='&#9733;';
else h+='<span class="empty-star">&#9733;</span>';
}
h+='</span>';
return h;
}

// ─── Loading ──────────────────────────────────────────────────────

async function load(){
try{
var resps=await Promise.all([
fetch(A+'/recipes').then(function(r){return r.json()}),
fetch(A+'/stats').then(function(r){return r.json()})
]);
recipes=resps[0].recipes||[];
renderStats(resps[1]||{});

try{
var ex=await fetch(A+'/extras/'+RESOURCE).then(function(r){return r.json()});
recipeExtras=ex||{};
recipes.forEach(function(r){
var x=recipeExtras[r.id];
if(!x)return;
Object.keys(x).forEach(function(k){if(r[k]===undefined)r[k]=x[k]});
});
}catch(e){recipeExtras={}}

populateCategoryFilter();
}catch(e){
console.error('load failed',e);
recipes=[];
}
render();
}

function populateCategoryFilter(){
var sel=document.getElementById('category-filter');
if(!sel)return;
var current=sel.value;
var seen={};
var cats=[];
recipes.forEach(function(r){
if(r.category&&!seen[r.category]){seen[r.category]=true;cats.push(r.category)}
});
cats.sort();
sel.innerHTML='<option value="">All Categories</option>'+cats.map(function(c){return'<option value="'+esc(c)+'"'+(c===current?' selected':'')+'>'+esc(c)+'</option>'}).join('');
}

function renderStats(s){
var total=s.total||0;
var cats=s.categories||0;
var avg=s.avg_rating||0;
var fav=s.favorites||0;
document.getElementById('stats').innerHTML=
'<div class="st"><div class="st-v">'+total+'</div><div class="st-l">Recipes</div></div>'+
'<div class="st"><div class="st-v">'+cats+'</div><div class="st-l">Categories</div></div>'+
'<div class="st"><div class="st-v">'+avg.toFixed(1)+'</div><div class="st-l">Avg Rating</div></div>'+
'<div class="st"><div class="st-v">'+fav+'</div><div class="st-l">Favorites</div></div>';
}

function render(){
var q=(document.getElementById('search').value||'').toLowerCase();
var cf=document.getElementById('category-filter').value;
var rf=parseInt(document.getElementById('rating-filter').value,10);

var f=recipes;
if(q)f=f.filter(function(r){
return(r.title||'').toLowerCase().includes(q)||
(r.ingredients||'').toLowerCase().includes(q)||
(r.tags||'').toLowerCase().includes(q);
});
if(cf)f=f.filter(function(r){return r.category===cf});
if(rf>0)f=f.filter(function(r){return(r.rating||0)>=rf});

if(!f.length){
var msg=window._emptyMsg||'No recipes yet. Add your first one.';
document.getElementById('grid').innerHTML='<div class="empty" style="grid-column:1/-1">'+esc(msg)+'</div>';
return;
}

var h='';
f.forEach(function(r){h+=cardHTML(r)});
document.getElementById('grid').innerHTML=h;
}

function cardHTML(r){
var h='<div class="card" onclick="openEdit(\''+esc(r.id)+'\')">';
h+='<div class="card-title">'+esc(r.title)+'</div>';
if(r.category)h+='<div class="card-cat">'+esc(r.category)+'</div>';
if(r.rating>0)h+=starsHTML(r.rating);
h+='<div class="card-meta">';
if(r.prep_time)h+='<span>Prep '+r.prep_time+'m</span>';
if(r.cook_time)h+='<span>Cook '+r.cook_time+'m</span>';
if(r.servings)h+='<span>Serves '+r.servings+'</span>';
h+='</div>';
if(r.tags){
var tagList=String(r.tags).split(',').map(function(t){return t.trim()}).filter(function(t){return t});
if(tagList.length){
h+='<div class="card-tags">';
tagList.forEach(function(t){h+='<span class="tag">#'+esc(t)+'</span>'});
h+='</div>';
}
}

// Custom fields
var customRows='';
fields.forEach(function(f){
if(!f.isCustom)return;
var v=r[f.name];
if(v===undefined||v===null||v==='')return;
customRows+='<div class="card-extra-row">';
customRows+='<span class="card-extra-label">'+esc(f.label)+'</span>';
customRows+='<span class="card-extra-val">'+esc(String(v))+'</span>';
customRows+='</div>';
});
if(customRows)h+='<div class="card-extra">'+customRows+'</div>';

h+='</div>';
return h;
}

// ─── Modal ────────────────────────────────────────────────────────

function fieldHTML(f,value){
var v=value;
if(v===undefined||v===null)v='';
var req=f.required?' *':'';
var ph=f.placeholder?(' placeholder="'+esc(f.placeholder)+'"'):'';
var h='<div class="fr"><label>'+esc(f.label)+req+'</label>';

if(f.type==='select'){
h+='<select id="f-'+f.name+'">';
if(!f.required)h+='<option value="">Select...</option>';
(f.options||[]).forEach(function(o){
var sel=(String(v)===String(o))?' selected':'';
h+='<option value="'+esc(String(o))+'"'+sel+'>'+esc(String(o))+'</option>';
});
h+='</select>';
}else if(f.type==='select_or_text'){
h+='<input list="dl-'+f.name+'" type="text" id="f-'+f.name+'" value="'+esc(String(v))+'"'+ph+'>';
h+='<datalist id="dl-'+f.name+'">';
var opts=(f.options||[]).slice();
recipes.forEach(function(rd){
if(rd.category&&opts.indexOf(rd.category)===-1)opts.push(rd.category);
});
opts.forEach(function(o){h+='<option value="'+esc(String(o))+'">'});
h+='</datalist>';
}else if(f.type==='textarea'){
var rows=f.name==='instructions'?5:3;
h+='<textarea id="f-'+f.name+'" rows="'+rows+'"'+ph+'>'+esc(String(v))+'</textarea>';
}else if(f.type==='rating'){
h+='<div class="rating-input" id="f-rating">';
for(var i=1;i<=5;i++){
var lit=i<=(v||0)?' lit':'';
h+='<button type="button" class="star-btn'+lit+'" data-val="'+i+'" onclick="setRating('+i+')">&#9733;</button>';
}
h+='<span class="star-clear" onclick="setRating(0)">clear</span>';
h+='</div>';
}else if(f.type==='number'||f.type==='integer'){
h+='<input type="number" id="f-'+f.name+'" value="'+esc(String(v))+'"'+ph+'>';
}else{
var inputType=f.type||'text';
h+='<input type="'+esc(inputType)+'" id="f-'+f.name+'" value="'+esc(String(v))+'"'+ph+'>';
}
h+='</div>';
return h;
}

function setRating(n){
currentRating=n;
var btns=document.querySelectorAll('.star-btn');
btns.forEach(function(b){
var v=parseInt(b.getAttribute('data-val'),10);
b.classList.toggle('lit',v<=n);
});
}

function formHTML(recipe){
var r=recipe||{};
var isEdit=!!recipe;
currentRating=r.rating||0;
var h='<h2>'+(isEdit?'Edit Recipe':'New Recipe')+'</h2>';

h+=fieldHTML(fieldByName('title'),r.title);
h+='<div class="row2">'+fieldHTML(fieldByName('category'),r.category)+fieldHTML(fieldByName('rating'),r.rating)+'</div>';
h+=fieldHTML(fieldByName('ingredients'),r.ingredients);
h+=fieldHTML(fieldByName('instructions'),r.instructions);
h+='<div class="row3">'+fieldHTML(fieldByName('prep_time'),r.prep_time)+fieldHTML(fieldByName('cook_time'),r.cook_time)+fieldHTML(fieldByName('servings'),r.servings)+'</div>';
h+=fieldHTML(fieldByName('tags'),r.tags);

var customFields=fields.filter(function(f){return f.isCustom});
if(customFields.length){
var label=window._customSectionLabel||'Additional Details';
h+='<div class="fr-section"><div class="fr-section-label">'+esc(label)+'</div>';
customFields.forEach(function(f){h+=fieldHTML(f,r[f.name])});
h+='</div>';
}

h+='<div class="acts">';
if(isEdit){
h+='<button class="btn btn-del" onclick="delRecipe()">Delete</button>';
}
h+='<button class="btn" onclick="closeModal()">Cancel</button>';
h+='<button class="btn btn-p" onclick="submit()">'+(isEdit?'Save':'Create')+'</button>';
h+='</div>';
return h;
}

function openNew(){
editId=null;
document.getElementById('mdl').innerHTML=formHTML();
document.getElementById('mbg').classList.add('open');
var t=document.getElementById('f-title');
if(t)t.focus();
}

function openEdit(id){
var r=null;
for(var i=0;i<recipes.length;i++){if(recipes[i].id===id){r=recipes[i];break}}
if(!r)return;
editId=id;
document.getElementById('mdl').innerHTML=formHTML(r);
document.getElementById('mbg').classList.add('open');
}

function closeModal(){
document.getElementById('mbg').classList.remove('open');
editId=null;
currentRating=0;
}

async function submit(){
var titleEl=document.getElementById('f-title');
if(!titleEl||!titleEl.value.trim()){alert('Title is required');return}

var body={};
var extras={};
fields.forEach(function(f){
if(f.type==='rating'){
if(f.isCustom)extras[f.name]=currentRating;
else body[f.name]=currentRating;
return;
}
var el=document.getElementById('f-'+f.name);
if(!el)return;
var val;
if(f.type==='number'||f.type==='integer')val=parseFloat(el.value)||0;
else val=el.value.trim();
if(f.isCustom)extras[f.name]=val;
else body[f.name]=val;
});

var savedId=editId;
try{
if(editId){
var r1=await fetch(A+'/recipes/'+editId,{method:'PUT',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});
if(!r1.ok){var e1=await r1.json().catch(function(){return{}});alert(e1.error||'Save failed');return}
}else{
var r2=await fetch(A+'/recipes',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});
if(!r2.ok){var e2=await r2.json().catch(function(){return{}});alert(e2.error||'Create failed');return}
var created=await r2.json();
savedId=created.id;
}
if(savedId&&Object.keys(extras).length){
await fetch(A+'/extras/'+RESOURCE+'/'+savedId,{method:'PUT',headers:{'Content-Type':'application/json'},body:JSON.stringify(extras)}).catch(function(){});
}
}catch(e){
alert('Network error: '+e.message);
return;
}
closeModal();
load();
}

async function delRecipe(){
if(!editId)return;
if(!confirm('Delete this recipe?'))return;
await fetch(A+'/recipes/'+editId,{method:'DELETE'});
closeModal();
load();
}

function esc(s){
if(s===undefined||s===null)return'';
var d=document.createElement('div');
d.textContent=String(s);
return d.innerHTML;
}

document.addEventListener('keydown',function(e){if(e.key==='Escape')closeModal()});

// ─── Personalization ──────────────────────────────────────────────

(function loadPersonalization(){
fetch('/api/config').then(function(r){return r.json()}).then(function(cfg){
if(!cfg||typeof cfg!=='object')return;

if(cfg.dashboard_title){
var h1=document.getElementById('dash-title');
if(h1)h1.innerHTML='<span>&#9670;</span> '+esc(cfg.dashboard_title);
document.title=cfg.dashboard_title;
}

if(cfg.empty_state_message)window._emptyMsg=cfg.empty_state_message;
if(cfg.primary_label)window._customSectionLabel=cfg.primary_label+' Details';

if(Array.isArray(cfg.categories)){
var catField=fieldByName('category');
if(catField)catField.options=cfg.categories;
}

if(Array.isArray(cfg.custom_fields)){
cfg.custom_fields.forEach(function(cf){
if(!cf||!cf.name||!cf.label)return;
if(fieldByName(cf.name))return;
fields.push({
name:cf.name,
label:cf.label,
type:cf.type||'text',
options:cf.options||[],
isCustom:true
});
});
}
}).catch(function(){
}).finally(function(){
load();
});
})();
</script>
</body>
</html>`
