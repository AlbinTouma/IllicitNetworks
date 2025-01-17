//This script generates latest stories. I'll update these in a simple js object and render them as <p> for now but will use CMS later.


const news = [

`
LYON, France – INTERPOL has published its first-ever <a href="https://www.interpol.int/News-and-Events/News/2025/INTERPOL-publishes-first-Silver-Notice-targeting-criminal-assets"> Silver Notice </a> to help trace and recover criminal assets, combat transnational organized crime and enhance international police cooperation.

The Notice, requested by Italy, seeks information on the assets belonging to a senior member of the mafia.
`,

`
A global law enforcement operation headed by <a href="https://thehackernews.com/2024/12/interpol-arrests-5500-in-global.html"> Interpol </a> has led to the arrest of more than 5,500 suspects involved in financial crimes and the seizure of more than $400 million in virtual assets and government-backed currencies.


`,

`
Egypt <a href="https://www.middleeasteye.net/news/egypt-sudan-smugglers-border-battles"> deploys military </a> to its southern border to stop smugglers from Sudan.
`,
`
Iran is using Swedish criminal networks to target Jews, <a href="https://sakerhetspolisen.se/ovriga-sidor/other-languages/english-engelska/press-room/news/news/2024-05-30-iran-is-using-criminal-networks-in-sweden.html">Sweden's security service warns</a>. 
The Iranian government had been using criminal networks within Sweden to carry out violent acts against other states, groups and individuals.
`,
`
THE US GOVERNMENT has branded rioting which broke out on the streets of Dublin in 2023 as “white identity terrorism”.
In <a href="https://www.state.gov/reports/country-reports-on-terrorism-2023/"> Country Reports on Terrorism 2023</a> released by the US Department of State, the violence was blamed on “Irish white supremacists and ultranationalists” spreading “anti-immigrant” disinformation online.

`

]


div = document.querySelector('.news')

news.forEach(item  => {
  console.log(item)
  p = document.createElement('p')
  p.innerHTML = item
  div.appendChild(p)
})

