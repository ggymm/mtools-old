const languages = [
  'bat',
  'c',
  'clojure',
  'coffeescript',
  'cpp',
  'csharp',
  'csp',
  'css',
  'dockerfile',
  'fsharp',
  'go',
  'graphql',
  'handlebars',
  'html',
  'ini',
  'java',
  'javascript',
  'json',
  'kotlin',
  'less',
  'lua',
  'markdown',
  'mysql',
  'objective-c',
  'pascal',
  'perl',
  'pgsql',
  'php',
  'plaintext',
  'powershell',
  'pug',
  'python',
  'r',
  'razor',
  'redis',
  'redshift',
  'ruby',
  'rust',
  'sb',
  'scheme',
  'scss',
  'shell',
  'sol',
  'sql',
  'st',
  'swift',
  'tcl',
  'typescript',
  'vb',
  'xml',
  'yaml'
]

const languageOptions = () => {
  const options = []
  for (const k in languages) {
    options.push({
      label: languages[k],
      value: languages[k]
    })
  }
  return options
}

export default languageOptions()
