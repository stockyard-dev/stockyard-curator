package server
import "net/http"
func(s *Server)dashboard(w http.ResponseWriter,r *http.Request){w.Header().Set("Content-Type","text/html");w.Write([]byte(dashHTML))}
const dashHTML=`<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><title>Curator</title>
<style>:root{--bg:#1a1410;--bg2:#241e18;--bg3:#2e261e;--rust:#e8753a;--leather:#a0845c;--cream:#f0e6d3;--cd:#bfb5a3;--cm:#7a7060;--gold:#d4a843;--green:#4a9e5c;--mono:'JetBrains Mono',monospace;--serif:'Libre Baskerville',serif}
*{margin:0;padding:0;box-sizing:border-box}body{background:var(--bg);color:var(--cream);font-family:var(--serif);line-height:1.6}
.hdr{padding:1rem 1.5rem;border-bottom:1px solid var(--bg3);display:flex;justify-content:space-between;align-items:center}.hdr h1{font-family:var(--mono);font-size:.9rem;letter-spacing:2px}
.main{padding:1.5rem;max-width:900px;margin:0 auto}
.search{width:100%;padding:.5rem .8rem;background:var(--bg2);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.78rem;margin-bottom:.8rem}
.cat-bar{display:flex;gap:.3rem;margin-bottom:1rem;flex-wrap:wrap}
.cat-btn{font-family:var(--mono);font-size:.6rem;padding:.2rem .5rem;border:1px solid var(--bg3);background:var(--bg);color:var(--cm);cursor:pointer}.cat-btn:hover{border-color:var(--leather)}.cat-btn.active{border-color:var(--rust);color:var(--rust)}
.grid{display:grid;grid-template-columns:repeat(auto-fill,minmax(260px,1fr));gap:.8rem}
.recipe{background:var(--bg2);border:1px solid var(--bg3);padding:1rem;cursor:pointer;transition:border-color .15s}
.recipe:hover{border-color:var(--leather)}
.recipe-title{font-size:1rem;margin-bottom:.3rem}
.recipe-meta{font-family:var(--mono);font-size:.6rem;color:var(--cm);display:flex;gap:.8rem;margin-bottom:.3rem}
.recipe-cat{font-family:var(--mono);font-size:.55rem;padding:.1rem .3rem;background:var(--bg3);color:var(--cm)}
.recipe-rating{color:var(--gold)}
.stars{letter-spacing:2px}
.btn{font-family:var(--mono);font-size:.6rem;padding:.25rem .6rem;cursor:pointer;border:1px solid var(--bg3);background:var(--bg);color:var(--cd)}.btn:hover{border-color:var(--leather);color:var(--cream)}
.btn-p{background:var(--rust);border-color:var(--rust);color:var(--bg)}
.modal-bg{display:none;position:fixed;inset:0;background:rgba(0,0,0,.6);z-index:100;align-items:center;justify-content:center}.modal-bg.open{display:flex}
.modal{background:var(--bg2);border:1px solid var(--bg3);padding:1.5rem;width:500px;max-width:90vw;max-height:90vh;overflow-y:auto}
.modal h2{font-family:var(--mono);font-size:.8rem;margin-bottom:1rem;color:var(--rust)}
.fr{margin-bottom:.5rem}.fr label{display:block;font-family:var(--mono);font-size:.55rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-bottom:.15rem}
.fr input,.fr select,.fr textarea{width:100%;padding:.35rem .5rem;background:var(--bg);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}
.fr textarea{min-height:80px;font-family:var(--serif);font-size:.85rem;line-height:1.7}
.acts{display:flex;gap:.4rem;justify-content:flex-end;margin-top:.8rem}
.detail{font-size:.9rem;color:var(--cd)}
.detail h3{font-family:var(--mono);font-size:.65rem;color:var(--leather);text-transform:uppercase;letter-spacing:1px;margin:1rem 0 .3rem}
.detail .ingredients{list-style:none;font-family:var(--mono);font-size:.78rem}
.detail .ingredients li{padding:.2rem 0;border-bottom:1px solid var(--bg3)}
.detail .instructions{white-space:pre-wrap;line-height:1.8}
.empty{text-align:center;padding:3rem;color:var(--cm);font-style:italic;font-size:.8rem}
</style></head><body>
<div class="hdr"><h1>CURATOR</h1><button class="btn btn-p" onclick="openForm()">+ Add Recipe</button></div>
<div class="main">
<input class="search" id="search" placeholder="Search recipes..." oninput="render()">
<div class="cat-bar" id="cats"></div>
<div class="grid" id="grid"></div>
</div>
<div class="modal-bg" id="mbg" onclick="if(event.target===this)cm()"><div class="modal" id="mdl"></div></div>
<script>
const A='/api';let recipes=[],filterCat='';
async function load(){const r=await fetch(A+'/recipes').then(r=>r.json());recipes=r.recipes||[];
const cats=[...new Set(recipes.map(r=>r.category).filter(c=>c))];
let h='<button class="cat-btn'+(filterCat===''?' active':'')+'" onclick="setCat(\'\')">All ('+recipes.length+')</button>';
cats.forEach(c=>{h+='<button class="cat-btn'+(filterCat===c?' active':'')+'" onclick="setCat(\''+c+'\')">'+esc(c)+'</button>';});
document.getElementById('cats').innerHTML=h;render();}
function setCat(c){filterCat=c;load();}
function render(){const q=(document.getElementById('search').value||'').toLowerCase();
let filtered=recipes.filter(r=>{if(filterCat&&r.category!==filterCat)return false;if(q&&!(r.title+r.category+r.tags).toLowerCase().includes(q))return false;return true;});
if(!filtered.length){document.getElementById('grid').innerHTML='<div class="empty">No recipes yet. Add your first one.</div>';return;}
let h='';filtered.forEach(r=>{
const stars='★'.repeat(r.rating)+'☆'.repeat(5-r.rating);
h+='<div class="recipe" onclick="viewRecipe(\''+r.id+'\')"><div class="recipe-title">'+esc(r.title)+'</div><div class="recipe-meta">';
if(r.prep_time)h+='<span>Prep: '+r.prep_time+'m</span>';
if(r.cook_time)h+='<span>Cook: '+r.cook_time+'m</span>';
h+='<span>Serves '+r.servings+'</span></div>';
if(r.category)h+='<span class="recipe-cat">'+esc(r.category)+'</span> ';
if(r.rating)h+='<span class="recipe-rating stars">'+stars+'</span>';
h+='</div>';});
document.getElementById('grid').innerHTML=h;}
function viewRecipe(id){const r=recipes.find(x=>x.id===id);if(!r)return;
let ingredients;try{ingredients=JSON.parse(r.ingredients||'[]')}catch(e){ingredients=r.ingredients?r.ingredients.split('\n'):[]}
let h='<h2>'+esc(r.title)+'</h2><div class="detail">';
h+='<div style="font-family:var(--mono);font-size:.65rem;color:var(--cm);margin-bottom:1rem">';
if(r.prep_time)h+='Prep: '+r.prep_time+'m · ';if(r.cook_time)h+='Cook: '+r.cook_time+'m · ';h+='Serves '+r.servings;
if(r.rating)h+=' · <span class="recipe-rating stars">'+'★'.repeat(r.rating)+'☆'.repeat(5-r.rating)+'</span>';
h+='</div>';
if(ingredients.length){h+='<h3>Ingredients</h3><ul class="ingredients">';ingredients.forEach(i=>{if(typeof i==="string"&&i.trim())h+='<li>'+esc(i)+'</li>';else if(i.name)h+='<li>'+esc(i.amount?i.amount+' ':'')+ esc(i.name)+'</li>';});h+='</ul>';}
if(r.instructions)h+='<h3>Instructions</h3><div class="instructions">'+esc(r.instructions)+'</div>';
h+='</div><div class="acts"><button class="btn" onclick="del(\''+r.id+'\')" style="color:var(--rust)">Delete</button><button class="btn" onclick="cm()">Close</button></div>';
document.getElementById('mdl').innerHTML=h;document.getElementById('mbg').classList.add('open');}
async function del(id){if(confirm('Delete?')){await fetch(A+'/recipes/'+id,{method:'DELETE'});cm();load();}}
function openForm(){document.getElementById('mdl').innerHTML='<h2>Add Recipe</h2><div class="fr"><label>Title</label><input id="f-t" placeholder="e.g. Pasta Carbonara"></div><div class="fr"><label>Ingredients (one per line)</label><textarea id="f-i" placeholder="200g spaghetti\n4 egg yolks\n100g pancetta"></textarea></div><div class="fr"><label>Instructions</label><textarea id="f-ins" rows="6" placeholder="Step-by-step..."></textarea></div><div style="display:grid;grid-template-columns:1fr 1fr 1fr;gap:.5rem"><div class="fr"><label>Prep (min)</label><input id="f-p" type="number" value="15"></div><div class="fr"><label>Cook (min)</label><input id="f-c" type="number" value="20"></div><div class="fr"><label>Servings</label><input id="f-s" type="number" value="4"></div></div><div class="fr"><label>Category</label><input id="f-cat" placeholder="e.g. Italian, Dessert, Salad"></div><div class="fr"><label>Rating (1-5)</label><input id="f-r" type="number" min="0" max="5" value="0"></div><div class="acts"><button class="btn" onclick="cm()">Cancel</button><button class="btn btn-p" onclick="sub()">Save</button></div>';document.getElementById('mbg').classList.add('open');}
async function sub(){const ings=document.getElementById('f-i').value.split('\n').filter(l=>l.trim());await fetch(A+'/recipes',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({title:document.getElementById('f-t').value,ingredients:JSON.stringify(ings),instructions:document.getElementById('f-ins').value,prep_time:parseInt(document.getElementById('f-p').value)||0,cook_time:parseInt(document.getElementById('f-c').value)||0,servings:parseInt(document.getElementById('f-s').value)||4,category:document.getElementById('f-cat').value,rating:parseInt(document.getElementById('f-r').value)||0})});cm();load();}
function cm(){document.getElementById('mbg').classList.remove('open');}
function esc(s){if(!s)return'';const d=document.createElement('div');d.textContent=s;return d.innerHTML;}
load();
</script></body></html>`
