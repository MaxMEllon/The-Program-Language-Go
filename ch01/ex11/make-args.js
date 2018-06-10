result = []
for (elem of document.querySelectorAll('.td.DescriptionCell > p > a')) {
  result.push(elem.href.replace('https://www.alexa.com/siteinfo/', ''))
}
result.join(' ')
