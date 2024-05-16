import{_ as t,r as e,o as p,c as o,a as n,b as c,F as l,e as r,d as s}from"./app.aa4fcc9f.js";const u={},i=r(`<h1 id="request-airdrop" tabindex="-1"><a class="header-anchor" href="#request-airdrop" aria-hidden="true">#</a> Request Airdrop</h1><p>Request some airdrop for testing.</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>
	sig<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">RequestAirdrop</span><span class="token punctuation">(</span>
		context<span class="token punctuation">.</span><span class="token function">TODO</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token string">&quot;9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an&quot;</span><span class="token punctuation">,</span> <span class="token comment">// address</span>
		<span class="token number">1e9</span><span class="token punctuation">,</span> <span class="token comment">// lamports (1 SOL = 10^9 lamports)</span>
	<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to request airdrop, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>sig<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br></div></div>`,3),k={class:"custom-container tip"},b=n("p",{class:"custom-container-title"},"TIP",-1),m=s("you can look up this tx on "),d={href:"https://explorer.solana.com/?cluster=devnet",target:"_blank",rel:"noopener noreferrer"},g=s("https://explorer.solana.com/?cluster=devnet");function _(f,q){const a=e("ExternalLinkIcon");return p(),o(l,null,[i,n("div",k,[b,n("p",null,[m,n("a",d,[g,c(a)])])])],64)}var x=t(u,[["render",_]]);export{x as default};
