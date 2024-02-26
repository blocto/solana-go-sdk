import{_ as n,e as s}from"./app.aa4fcc9f.js";const a={},t=s(`<h1 id="get-signatures-for-address" tabindex="-1"><a class="header-anchor" href="#get-signatures-for-address" aria-hidden="true">#</a> Get Signatures For Address</h1><p>Fetch tx histroy.</p><h2 id="all" tabindex="-1"><a class="header-anchor" href="#all" aria-hidden="true">#</a> All</h2><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>	<span class="token comment">// get all (limit is between 1 ~ 1,000, default is 1,000)</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddress</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> target<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br></div></div><h2 id="limit" tabindex="-1"><a class="header-anchor" href="#limit" aria-hidden="true">#</a> Limit</h2><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>	<span class="token comment">// get latest X tx</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddressWithConfig</span><span class="token punctuation">(</span>
			context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
			target<span class="token punctuation">,</span>
			rpc<span class="token punctuation">.</span>GetSignaturesForAddressConfig<span class="token punctuation">{</span>
				Limit<span class="token punctuation">:</span> <span class="token number">5</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br></div></div><h2 id="range" tabindex="-1"><a class="header-anchor" href="#range" aria-hidden="true">#</a> Range</h2><p>context:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>	<span class="token comment">/*
		if a txhash list like:

		(new)
		3gwJwVorVprqZmm1ULAp9bKQy6sQ7skG11XYdMSQigA936MBANcSBy6NcNJF2yPf9ycgzZ6vFd4pjAY7qSko61Au
		wTEnw3vpBthzLUD6gv9B3aC4dQNp4hews85ipM3w9MGZAh38HZ2im9LaWY7aRusVN5Wj33mNvqSRDNyC43u6GQs
		3e6dRv5KnvpU43VjVjbsubvPR1yFK9b922WcTugyTBSdWdToeCK16NccSaxY6XJ5yi51UswP3ZDe3VJBZTVg2MCW
		2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy
		2uFaNDgQWZsgZvR6s3WQKwaCxFgS4ML7xrZyAqgmuTSEuGmrWyCcTrjtajr6baYR6FaVLZ4PWgyt55EmTcT8S7Sg
		4XGVHHpLW99AUFEd6RivasG57vqu4EMMNdcQdmphepmW484dMYtWLkYw4nSNnSpKiDoYDbSu9ksxECNKBk2JEyHQ
		3kjLJokcYqAhQjERCVutv5gdUuQ1HsxSCcFsJdQbqNkqd5ML8WRaZJguZgpWH8isCfyEN8YktxxPPNJURhAtvUKE
		(old)
	*/</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br></div></div><h3 id="before" tabindex="-1"><a class="header-anchor" href="#before" aria-hidden="true">#</a> Before</h3><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>	<span class="token comment">// you can fetch the last 3 tx by</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddressWithConfig</span><span class="token punctuation">(</span>
			context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
			target<span class="token punctuation">,</span>
			rpc<span class="token punctuation">.</span>GetSignaturesForAddressConfig<span class="token punctuation">{</span>
				Before<span class="token punctuation">:</span> <span class="token string">&quot;2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy&quot;</span><span class="token punctuation">,</span>
				Limit<span class="token punctuation">:</span>  <span class="token number">3</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br></div></div><h3 id="until" tabindex="-1"><a class="header-anchor" href="#until" aria-hidden="true">#</a> Until</h3><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>	<span class="token comment">// you can fetch the latest 3 tx by \`until\`</span>
	<span class="token comment">// * the result will be different if there are some newer txs added.</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddressWithConfig</span><span class="token punctuation">(</span>
			context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
			target<span class="token punctuation">,</span>
			rpc<span class="token punctuation">.</span>GetSignaturesForAddressConfig<span class="token punctuation">{</span>
				Until<span class="token punctuation">:</span> <span class="token string">&quot;2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy&quot;</span><span class="token punctuation">,</span>
				Limit<span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br></div></div><h2 id="full-code" tabindex="-1"><a class="header-anchor" href="#full-code" aria-hidden="true">#</a> Full Code</h2><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
	<span class="token string">&quot;github.com/davecgh/go-spew/spew&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>
	target <span class="token operator">:=</span> <span class="token string">&quot;Memo1UhkJRfHyvLMcVucJwxXeuD728EqVDDwQDxFMNo&quot;</span>

	<span class="token comment">// get all (limit is between 1 ~ 1,000, default is 1,000)</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddress</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> target<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">// get latest X tx</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddressWithConfig</span><span class="token punctuation">(</span>
			context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
			target<span class="token punctuation">,</span>
			rpc<span class="token punctuation">.</span>GetSignaturesForAddressConfig<span class="token punctuation">{</span>
				Limit<span class="token punctuation">:</span> <span class="token number">5</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">/*
		if a txhash list like:

		(new)
		3gwJwVorVprqZmm1ULAp9bKQy6sQ7skG11XYdMSQigA936MBANcSBy6NcNJF2yPf9ycgzZ6vFd4pjAY7qSko61Au
		wTEnw3vpBthzLUD6gv9B3aC4dQNp4hews85ipM3w9MGZAh38HZ2im9LaWY7aRusVN5Wj33mNvqSRDNyC43u6GQs
		3e6dRv5KnvpU43VjVjbsubvPR1yFK9b922WcTugyTBSdWdToeCK16NccSaxY6XJ5yi51UswP3ZDe3VJBZTVg2MCW
		2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy
		2uFaNDgQWZsgZvR6s3WQKwaCxFgS4ML7xrZyAqgmuTSEuGmrWyCcTrjtajr6baYR6FaVLZ4PWgyt55EmTcT8S7Sg
		4XGVHHpLW99AUFEd6RivasG57vqu4EMMNdcQdmphepmW484dMYtWLkYw4nSNnSpKiDoYDbSu9ksxECNKBk2JEyHQ
		3kjLJokcYqAhQjERCVutv5gdUuQ1HsxSCcFsJdQbqNkqd5ML8WRaZJguZgpWH8isCfyEN8YktxxPPNJURhAtvUKE
		(old)
	*/</span>

	<span class="token comment">// you can fetch the last 3 tx by</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddressWithConfig</span><span class="token punctuation">(</span>
			context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
			target<span class="token punctuation">,</span>
			rpc<span class="token punctuation">.</span>GetSignaturesForAddressConfig<span class="token punctuation">{</span>
				Before<span class="token punctuation">:</span> <span class="token string">&quot;2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy&quot;</span><span class="token punctuation">,</span>
				Limit<span class="token punctuation">:</span>  <span class="token number">3</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">// you can fetch the latest 3 tx by \`until\`</span>
	<span class="token comment">// * the result will be different if there are some newer txs added.</span>
	<span class="token punctuation">{</span>
		res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetSignaturesForAddressWithConfig</span><span class="token punctuation">(</span>
			context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
			target<span class="token punctuation">,</span>
			rpc<span class="token punctuation">.</span>GetSignaturesForAddressConfig<span class="token punctuation">{</span>
				Until<span class="token punctuation">:</span> <span class="token string">&quot;2nYnHvbVuwmYeara3VjoCt9uS8ZXrSra5DRK7QBT8i5acoBiSK3FQY2vsaDSJQ6QX5i1pkvyRRjL1oUATMLZEsqy&quot;</span><span class="token punctuation">,</span>
				Limit<span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to GetSignaturesForAddress, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		spew<span class="token punctuation">.</span><span class="token function">Dump</span><span class="token punctuation">(</span>res<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br><span class="line-number">31</span><br><span class="line-number">32</span><br><span class="line-number">33</span><br><span class="line-number">34</span><br><span class="line-number">35</span><br><span class="line-number">36</span><br><span class="line-number">37</span><br><span class="line-number">38</span><br><span class="line-number">39</span><br><span class="line-number">40</span><br><span class="line-number">41</span><br><span class="line-number">42</span><br><span class="line-number">43</span><br><span class="line-number">44</span><br><span class="line-number">45</span><br><span class="line-number">46</span><br><span class="line-number">47</span><br><span class="line-number">48</span><br><span class="line-number">49</span><br><span class="line-number">50</span><br><span class="line-number">51</span><br><span class="line-number">52</span><br><span class="line-number">53</span><br><span class="line-number">54</span><br><span class="line-number">55</span><br><span class="line-number">56</span><br><span class="line-number">57</span><br><span class="line-number">58</span><br><span class="line-number">59</span><br><span class="line-number">60</span><br><span class="line-number">61</span><br><span class="line-number">62</span><br><span class="line-number">63</span><br><span class="line-number">64</span><br><span class="line-number">65</span><br><span class="line-number">66</span><br><span class="line-number">67</span><br><span class="line-number">68</span><br><span class="line-number">69</span><br><span class="line-number">70</span><br><span class="line-number">71</span><br><span class="line-number">72</span><br><span class="line-number">73</span><br><span class="line-number">74</span><br><span class="line-number">75</span><br><span class="line-number">76</span><br><span class="line-number">77</span><br><span class="line-number">78</span><br><span class="line-number">79</span><br><span class="line-number">80</span><br><span class="line-number">81</span><br><span class="line-number">82</span><br><span class="line-number">83</span><br><span class="line-number">84</span><br><span class="line-number">85</span><br><span class="line-number">86</span><br></div></div>`,15);function p(e,c){return t}var u=n(a,[["render",p]]);export{u as default};
