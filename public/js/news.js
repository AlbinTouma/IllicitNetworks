//This script generates latest stories. I'll update these in a simple js object and render them as <p> for now but will use CMS later.


const news = [
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

