const login = function login(){
  const url = '/login'
  const method = 'POST'
  const headers = {
      'Content-Type': 'application/json; charset=UTF-8'
  }



  const body = JSON.stringify({
      name: 'gaku0211', // 仮
      password: 'password',　// 仮
  })

  fetch(url, {method, headers, body}).then(response => {
      if(response.ok) {
          console.log('hoge');
          return response.json()
      } else {
          alert('Faild to login. Please retry')
          return {token: ''}
      }
  }).then(json => {
      const token = json.token
      console.log(token);
      if(token.length > 0) {
          localStorage.setItem('token', token)
          location.href = '/posts'
      }
  })
}