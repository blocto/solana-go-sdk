import{_ as n,e as s}from"./app.aa4fcc9f.js";const a={},p=s(`<h1 id="hello" tabindex="-1"><a class="header-anchor" href="#hello" aria-hidden="true">#</a> Hello</h1><h3 id="client" tabindex="-1"><a class="header-anchor" href="#client" aria-hidden="true">#</a> client</h3><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token comment">// an instruction is the smallest unit in a tx.</span>
<span class="token comment">// each instruction represent an action with a program.</span>
<span class="token comment">// there are three basic fields in an instruction:</span>
<span class="token comment">//   - program id</span>
<span class="token comment">//     the program which you would like to interact with</span>
<span class="token comment">//   - account meta list</span>
<span class="token comment">//     accounts which are used in the program</span>
<span class="token comment">//   - data</span>
<span class="token comment">//     a u8 array.</span>

<span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>

	<span class="token string">&quot;github.com/blocto/solana-go-sdk/client&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/common&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/rpc&quot;</span>
	<span class="token string">&quot;github.com/blocto/solana-go-sdk/types&quot;</span>
<span class="token punctuation">)</span>

<span class="token comment">// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz</span>
<span class="token keyword">var</span> feePayer<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">=</span> types<span class="token punctuation">.</span><span class="token function">AccountFromBase58</span><span class="token punctuation">(</span><span class="token string">&quot;4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN&quot;</span><span class="token punctuation">)</span>

<span class="token keyword">var</span> programId <span class="token operator">=</span> common<span class="token punctuation">.</span><span class="token function">PublicKeyFromString</span><span class="token punctuation">(</span><span class="token string">&quot;EGz5CDh7dG7BwzqL7y5ePpZNvrw7ehr4E4oGRhCCpiEK&quot;</span><span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c <span class="token operator">:=</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>rpc<span class="token punctuation">.</span>DevnetRPCEndpoint<span class="token punctuation">)</span>

	res<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">GetLatestBlockhash</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to get latest blockhash, err: %v\\n&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">// our first program won&#39;t use any accounts and parse any data. we leave them empty atm.</span>
	tx<span class="token punctuation">,</span> err <span class="token operator">:=</span> types<span class="token punctuation">.</span><span class="token function">NewTransaction</span><span class="token punctuation">(</span>types<span class="token punctuation">.</span>NewTransactionParam<span class="token punctuation">{</span>
		Signers<span class="token punctuation">:</span> <span class="token punctuation">[</span><span class="token punctuation">]</span>types<span class="token punctuation">.</span>Account<span class="token punctuation">{</span>feePayer<span class="token punctuation">}</span><span class="token punctuation">,</span>
		Message<span class="token punctuation">:</span> types<span class="token punctuation">.</span><span class="token function">NewMessage</span><span class="token punctuation">(</span>types<span class="token punctuation">.</span>NewMessageParam<span class="token punctuation">{</span>
			FeePayer<span class="token punctuation">:</span>        feePayer<span class="token punctuation">.</span>PublicKey<span class="token punctuation">,</span>
			RecentBlockhash<span class="token punctuation">:</span> res<span class="token punctuation">.</span>Blockhash<span class="token punctuation">,</span>
			Instructions<span class="token punctuation">:</span> <span class="token punctuation">[</span><span class="token punctuation">]</span>types<span class="token punctuation">.</span>Instruction<span class="token punctuation">{</span>
				<span class="token punctuation">{</span>
					ProgramID<span class="token punctuation">:</span> programId<span class="token punctuation">,</span>
					Accounts<span class="token punctuation">:</span>  <span class="token punctuation">[</span><span class="token punctuation">]</span>types<span class="token punctuation">.</span>AccountMeta<span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
					Data<span class="token punctuation">:</span>      <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span>
				<span class="token punctuation">}</span><span class="token punctuation">,</span>
			<span class="token punctuation">}</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to new a tx, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	sig<span class="token punctuation">,</span> err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">SendTransaction</span><span class="token punctuation">(</span>context<span class="token punctuation">.</span><span class="token function">Background</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> tx<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalf</span><span class="token punctuation">(</span><span class="token string">&quot;failed to send the tx, err: %v&quot;</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token comment">// 5eCfcdKGzn8yLGfgX6V1AfjAavTQRGQ936Zyuqejf9W1tTsjABjHTrc66nv5g9qKStmRPr3FeCVznADuCnJ8Zfbq</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>sig<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br><span class="line-number">31</span><br><span class="line-number">32</span><br><span class="line-number">33</span><br><span class="line-number">34</span><br><span class="line-number">35</span><br><span class="line-number">36</span><br><span class="line-number">37</span><br><span class="line-number">38</span><br><span class="line-number">39</span><br><span class="line-number">40</span><br><span class="line-number">41</span><br><span class="line-number">42</span><br><span class="line-number">43</span><br><span class="line-number">44</span><br><span class="line-number">45</span><br><span class="line-number">46</span><br><span class="line-number">47</span><br><span class="line-number">48</span><br><span class="line-number">49</span><br><span class="line-number">50</span><br><span class="line-number">51</span><br><span class="line-number">52</span><br><span class="line-number">53</span><br><span class="line-number">54</span><br><span class="line-number">55</span><br><span class="line-number">56</span><br><span class="line-number">57</span><br><span class="line-number">58</span><br><span class="line-number">59</span><br><span class="line-number">60</span><br><span class="line-number">61</span><br><span class="line-number">62</span><br><span class="line-number">63</span><br></div></div><h3 id="program" tabindex="-1"><a class="header-anchor" href="#program" aria-hidden="true">#</a> program</h3><div class="language-rust ext-rs line-numbers-mode"><pre class="language-rust"><code><span class="token keyword">use</span> <span class="token namespace">solana_program<span class="token punctuation">::</span></span><span class="token punctuation">{</span>
    <span class="token namespace">account_info<span class="token punctuation">::</span></span><span class="token class-name">AccountInfo</span><span class="token punctuation">,</span> entrypoint<span class="token punctuation">,</span> <span class="token namespace">entrypoint<span class="token punctuation">::</span></span><span class="token class-name">ProgramResult</span><span class="token punctuation">,</span> msg<span class="token punctuation">,</span> <span class="token namespace">pubkey<span class="token punctuation">::</span></span><span class="token class-name">Pubkey</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span><span class="token punctuation">;</span>

<span class="token macro property">entrypoint!</span><span class="token punctuation">(</span>process_instruction<span class="token punctuation">)</span><span class="token punctuation">;</span>

<span class="token keyword">fn</span> <span class="token function-definition function">process_instruction</span><span class="token punctuation">(</span>
    _program_id<span class="token punctuation">:</span> <span class="token operator">&amp;</span><span class="token class-name">Pubkey</span><span class="token punctuation">,</span>
    _accounts<span class="token punctuation">:</span> <span class="token operator">&amp;</span><span class="token punctuation">[</span><span class="token class-name">AccountInfo</span><span class="token punctuation">]</span><span class="token punctuation">,</span>
    _instruction_data<span class="token punctuation">:</span> <span class="token operator">&amp;</span><span class="token punctuation">[</span><span class="token keyword">u8</span><span class="token punctuation">]</span><span class="token punctuation">,</span>
<span class="token punctuation">)</span> <span class="token punctuation">-&gt;</span> <span class="token class-name">ProgramResult</span> <span class="token punctuation">{</span>
    <span class="token macro property">msg!</span><span class="token punctuation">(</span><span class="token string">&quot;hello&quot;</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
    <span class="token class-name">Ok</span><span class="token punctuation">(</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br></div></div>`,5);function t(e,o){return p}var u=n(a,[["render",t]]);export{u as default};
