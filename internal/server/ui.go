package server
import "net/http"
func(s *Server)dashboard(w http.ResponseWriter,r *http.Request){w.Header().Set("Content-Type","text/html");w.Write([]byte(dashHTML))}
const dashHTML=`<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><title>Curator</title>
<style>:root{--bg:#1a1410;--bg2:#241e18;--bg3:#2e261e;--rust:#e8753a;--leather:#a0845c;--cream:#f0e6d3;--cd:#bfb5a3;--cm:#7a7060;--gold:#d4a843;--green:#4a9e5c;--red:#c94444;--mono:'JetBrains Mono',monospace;--serif:'Libre Baskerville',serif}
*{margin:0;padding:0;box-sizing:border-box}body{background:var(--bg);color:var(--cream);font-family:var(--serif);line-height:1.6}
.hdr{padding:1rem 1.5rem;border-bottom:1px solid var(--bg3);display:flex;justify-content:space-between;align-items:center}.hdr h1{font-family:var(--mono);font-size:.9rem;letter-spacing:2px;color:var(--rust)}
.stats{display:grid;grid-template-columns:repeat(4,1fr);gap:.8rem;padding:1rem 1.5rem;max-width:900px;margin:0 auto}
.stat{background:var(--bg2);border:1px solid var(--bg3);padding:.8rem 1rem;text-align:center}
.stat-v{font-family:var(--mono);font-size:1.5rem;color:var(--cream)}.stat-l{font-family:var(--mono);font-size:.55rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-top:.1rem}
.main{padding:0 1.5rem 1.5rem;max-width:900px;margin:0 auto}
.bar{display:flex;gap:.5rem;margin-bottom:.8rem;align-items:center;flex-wrap:wrap}
.search{flex:1;min-width:180px;padding:.5rem .8rem;background:var(--bg2);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.78rem;outline:none}.search:focus{border-color:var(--leather)}
.cat-bar{display:flex;gap:.3rem;margin-bottom:1rem;flex-wrap:wrap}
.cat-btn{font-family:var(--mono);font-size:.6rem;padding:.2rem .5rem;border:1px solid var(--bg3);background:var(--bg);color:var(--cm);cursor:pointer;transition:all .15s}.cat-btn:hover{border-color:var(--leather);color:var(--cream)}.cat-btn.active{border-color:var(--rust);color:var(--rust);background:rgba(232,117,58,.08)}
.grid{display:grid;grid-template-columns:repeat(auto-fill,minmax(260px,1fr));gap:.8rem}
.card{background:var(--bg2);border:1px solid var(--bg3);padding:1rem;cursor:pointer;transition:border-color .15s}.card:hover{border-color:var(--leather)}
.card-title{font-size:.95rem;margin-bottom:.3rem;color:var(--cream)}
.card-meta{font-family:var(--mono);font-size:.6rem;color:var(--cm);display:flex;gap:.8rem;margin-bottom:.4rem;flex-wrap:wrap}
.card-foot{display:flex;gap:.4rem;align-items:center;flex-wrap:wrap}
.tag{font-family:var(--mono);font-size:.5rem;padding:.1rem .35rem;background:var(--bg3);color:var(--cm);display:inline-block}
.tag-cat{border-left:2px solid var(--leather)}
.stars{color:var(--gold);font-size:.7rem;letter-spacing:1px}
.btn{font-family:var(--mono);font-size:.6rem;padding:.3rem .7rem;cursor:pointer;border:1px solid var(--bg3);background:var(--bg);color:var(--cd);transition:all .15s}.btn:hover{border-color:var(--leather);color:var(--cream)}
.btn-p{background:var(--rust);border-color:var(--rust);color:var(--bg)}.btn-p:hover{background:#d06830}
.btn-d{color:var(--red);border-color:transparent}.btn-d:hover{border-color:var(--red)}
.modal-bg{display:none;position:fixed;inset:0;background:rgba(0,0,0,.65);z-index:100;align-items:center;justify-content:center}.modal-bg.open{display:flex}
.modal{background:var(--bg2);border:1px solid var(--bg3);padding:1.5rem;width:560px;max-width:92vw;max-height:90vh;overflow-y:auto}
.modal h2{font-family:var(--mono);font-size:.8rem;margin-bottom:1rem;color:var(--rust)}
.fr{margin-bottom:.6rem}.fr label{display:block;font-family:var(--mono);font-size:.55rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-bottom:.15rem}
.fr input,.fr select,.fr textarea{width:100%;padding:.4rem .5rem;background:var(--bg);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.72rem;outline:none}.fr input:focus,.fr textarea:focus{border-color:var(--leather)}
.fr textarea{min-height:80px;font-family:var(--serif);font-size:.82rem;line-height:1.7;resize:vertical}
.fr-row{display:grid;grid-template-columns:1fr 1fr 1fr;gap:.5rem}
.fr-stars{display:flex;gap:.3rem;margin-top:.2rem}
.fr-star{font-size:1.3rem;cursor:pointer;color:var(--bg3);transition:color .1s}.fr-star:hover,.fr-star.on{color:var(--gold)}
.acts{display:flex;gap:.4rem;justify-content:flex-end;margin-top:1rem}
.detail-head{display:flex;justify-content:space-between;align-items:flex-start;margin-bottom:.8rem}
.detail h3{font-family:var(--mono);font-size:.6rem;color:var(--leather);text-transform:uppercase;letter-spacing:1px;margin:1rem 0 .3rem}
.detail-meta{font-family:var(--mono);font-size:.62rem;color:var(--cm);display:flex;gap:1rem;flex-wrap:wrap;margin-bottom:.8rem}
.ing-list{list-style:none}.ing-list li{padding:.3rem 0;border-bottom:1px solid var(--bg3);font-family:var(--mono);font-size:.75rem;color:var(--cd)}
.ing-list li::before{content:'•';color:var(--rust);margin-right:.5rem}
.inst{white-space:pre-wrap;line-height:1.9;font-size:.88rem;color:var(--cd)}
.empty{text-align:center;padding:3rem;color:var(--cm);font-style:italic;font-size:.8rem}
.detail-tags{display:flex;gap:.3rem;flex-wrap:wrap;margin-top:.6rem}
@media(max-width:600px){.stats{grid-template-columns:repeat(2,1fr)}.fr-row{grid-template-columns:1fr}}
</style></head><body>
<div class="hdr"><h1>CURATOR</h1><button class="btn btn-p" onclick="openForm()">+ Add Recipe</button></div>
<div class="stats" id="stats"></div>
<div class="main">
<div class="bar"><input class="search" id="search" placeholder="Search recipes..." oninput="render()"></div>
<div class="cat-bar" id="cats"></div>
<div class="grid" id="grid"></div>
</div>
<div class="modal-bg" id="mbg" onclick="if(event.target===this)cm()"><div class="modal" id="mdl"></div></div>
<script>
const A='/api';let recipes=[],filterCat='',editId=null,formRating=0;
async function load(){
const[rr,ss]=await Promise.all([fetch(A+'/recipes').then(r=>r.json()),fetch(A+'/stats').then(r=>r.json())]);
recipes=rr.recipes||[];
document.getElementById('stats').innerHTML=
'<div class="stat"><div class="stat-v">'+ss.total+'</div><div class="stat-l">Recipes</div></div>'+
'<div class="stat"><div class="stat-v">'+ss.categories+'</div><div class="stat-l">Categories</div></div>'+
'<div class="stat"><div class="stat-v">'+(ss.avg_rating?Number(ss.avg_rating).toFixed(1):'—')+'</div><div class="stat-l">Avg Rating</div></div>'+
'<div class="stat"><div class="stat-v">'+ss.favorites+'</div><div class="stat-l">Favorites</div></div>';
const cats=[...new Set(recipes.map(r=>r.category).filter(c=>c))].sort();
let h='<button class="cat-btn'+(filterCat===''?' active':'')+'" onclick="setCat(\x27\x27)">All ('+recipes.length+')</button>';
cats.forEach(c=>{const n=recipes.filter(r=>r.category===c).length;h+='<button class="cat-btn'+(filterCat===c?' active':'')+'" onclick="setCat(\x27'+esc(c)+'\x27)">'+esc(c)+' ('+n+')</button>';});
document.getElementById('cats').innerHTML=h;render();}
function setCat(c){filterCat=c;
document.querySelectorAll('.cat-btn').forEach(b=>{b.classList.remove('active')});
event.target.classList.add('active');render();}
function render(){const q=(document.getElementById('search').value||'').toLowerCase();
let f=recipes.filter(r=>{if(filterCat&&r.category!==filterCat)return false;if(q&&!(r.title+' '+r.category+' '+r.tags).toLowerCase().includes(q))return false;return true;});
if(!f.length){document.getElementById('grid').innerHTML='<div class="empty">No recipes found. Add your first recipe to get started.</div>';return;}
let h='';f.forEach(r=>{
const stars=r.rating?'<span class="stars">'+'★'.repeat(r.rating)+'☆'.repeat(5-r.rating)+'</span>':'';
const tags=(r.tags||'').split(',').filter(t=>t.trim()).map(t=>'<span class="tag">'+esc(t.trim())+'</span>').join('');
let totalTime=(r.prep_time||0)+(r.cook_time||0);
h+='<div class="card" onclick="viewRecipe(\x27'+r.id+'\x27)"><div class="card-title">'+esc(r.title)+'</div><div class="card-meta">';
if(totalTime)h+='<span>'+totalTime+' min</span>';
h+='<span>Serves '+r.servings+'</span></div><div class="card-foot">';
if(r.category)h+='<span class="tag tag-cat">'+esc(r.category)+'</span>';
h+=tags+stars+'</div></div>';});
document.getElementById('grid').innerHTML=h;}
function viewRecipe(id){const r=recipes.find(x=>x.id===id);if(!r)return;
let ings;try{ings=JSON.parse(r.ingredients||'[]')}catch(e){ings=r.ingredients?r.ingredients.split('\n').filter(l=>l.trim()):[]}
const stars=r.rating?'<span class="stars">'+'★'.repeat(r.rating)+'☆'.repeat(5-r.rating)+'</span>':'';
const tags=(r.tags||'').split(',').filter(t=>t.trim());
let h='<div class="detail"><div class="detail-head"><h2 style="color:var(--cream);font-family:var(--serif);font-size:1.2rem">'+esc(r.title)+'</h2>'+stars+'</div>';
h+='<div class="detail-meta">';
if(r.prep_time)h+='<span>Prep '+r.prep_time+'m</span>';
if(r.cook_time)h+='<span>Cook '+r.cook_time+'m</span>';
let total=(r.prep_time||0)+(r.cook_time||0);if(total)h+='<span>Total '+total+'m</span>';
h+='<span>Serves '+r.servings+'</span>';
if(r.category)h+='<span class="tag tag-cat">'+esc(r.category)+'</span>';
h+='</div>';
if(ings.length){h+='<h3>Ingredients</h3><ul class="ing-list">';ings.forEach(i=>{if(typeof i==="string"&&i.trim())h+='<li>'+esc(i)+'</li>';else if(i&&i.name)h+='<li>'+esc(i.amount?i.amount+' ':'')+esc(i.name)+'</li>';});h+='</ul>';}
if(r.instructions)h+='<h3>Instructions</h3><div class="inst">'+esc(r.instructions)+'</div>';
if(tags.length){h+='<div class="detail-tags">';tags.forEach(t=>h+='<span class="tag">'+esc(t)+'</span>');h+='</div>';}
h+='</div><div class="acts"><button class="btn btn-d" onclick="del(\x27'+r.id+'\x27)">Delete</button><button class="btn" onclick="openForm(\x27'+r.id+'\x27)">Edit</button><button class="btn" onclick="cm()">Close</button></div>';
document.getElementById('mdl').innerHTML=h;document.getElementById('mbg').classList.add('open');}
async function del(id){if(confirm('Delete this recipe?')){await fetch(A+'/recipes/'+id,{method:'DELETE'});cm();load();}}
function openForm(id){editId=id||null;const r=id?recipes.find(x=>x.id===id):null;
formRating=r?r.rating:0;
let ings='';if(r){try{const a=JSON.parse(r.ingredients||'[]');ings=a.map(i=>typeof i==='string'?i:((i.amount||'')+' '+(i.name||'')).trim()).join('\n');}catch(e){ings=r.ingredients||'';}}
let h='<h2>'+(r?'Edit':'Add')+' Recipe</h2>';
h+='<div class="fr"><label>Title</label><input id="f-t" value="'+esc(r?r.title:'')+'" placeholder="e.g. Pasta Carbonara"></div>';
h+='<div class="fr"><label>Ingredients (one per line)</label><textarea id="f-i" placeholder="200g spaghetti\n4 egg yolks\n100g pancetta">'+esc(ings)+'</textarea></div>';
h+='<div class="fr"><label>Instructions</label><textarea id="f-ins" rows="6" placeholder="Step-by-step instructions...">'+esc(r?r.instructions:'')+'</textarea></div>';
h+='<div class="fr-row"><div class="fr"><label>Prep (min)</label><input id="f-p" type="number" value="'+(r?r.prep_time:15)+'"></div>';
h+='<div class="fr"><label>Cook (min)</label><input id="f-c" type="number" value="'+(r?r.cook_time:20)+'"></div>';
h+='<div class="fr"><label>Servings</label><input id="f-s" type="number" value="'+(r?r.servings:4)+'"></div></div>';
h+='<div class="fr"><label>Category</label><input id="f-cat" value="'+esc(r?r.category:'')+'" placeholder="e.g. Italian, Dessert, Salad"></div>';
h+='<div class="fr"><label>Tags (comma separated)</label><input id="f-tags" value="'+esc(r?r.tags:'')+'" placeholder="quick, weeknight, vegetarian"></div>';
h+='<div class="fr"><label>Rating</label><div class="fr-stars" id="f-stars">';
for(let i=1;i<=5;i++)h+='<span class="fr-star'+(formRating>=i?' on':'')+'" data-v="'+i+'" onclick="setRating('+i+')">★</span>';
h+='</div></div>';
h+='<div class="acts"><button class="btn" onclick="cm()">Cancel</button><button class="btn btn-p" onclick="sub()">Save</button></div>';
document.getElementById('mdl').innerHTML=h;document.getElementById('mbg').classList.add('open');}
function setRating(v){formRating=v;document.querySelectorAll('.fr-star').forEach(s=>{s.classList.toggle('on',parseInt(s.dataset.v)<=v)});}
async function sub(){const ings=document.getElementById('f-i').value.split('\n').filter(l=>l.trim());
const body={title:document.getElementById('f-t').value,ingredients:JSON.stringify(ings),instructions:document.getElementById('f-ins').value,prep_time:parseInt(document.getElementById('f-p').value)||0,cook_time:parseInt(document.getElementById('f-c').value)||0,servings:parseInt(document.getElementById('f-s').value)||4,category:document.getElementById('f-cat').value,tags:document.getElementById('f-tags').value,rating:formRating};
if(editId){await fetch(A+'/recipes/'+editId,{method:'PUT',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});}
else{await fetch(A+'/recipes',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});}
cm();load();}
function cm(){document.getElementById('mbg').classList.remove('open');editId=null;formRating=0;}
function esc(s){if(!s)return'';const d=document.createElement('div');d.textContent=String(s);return d.innerHTML;}
load();
</script></body></html>`
