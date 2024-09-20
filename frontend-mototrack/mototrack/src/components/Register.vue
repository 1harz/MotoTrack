<template>
  <div class="background">
    <div class="register-container">
      <h2>Register</h2>
      <form @submit.prevent="register">
        <div class="form-group">
          <input type="email" v-model="email" class="form-control" placeholder="Email" required>
        </div>
        <div class="form-group">
          <input type="password" v-model="password" class="form-control" placeholder="Password" required>
        </div>
        <button type="submit" class="btn btn-dark btn-block">Sign Up</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      email: '',
      password: '',
      errorMessage: ''
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post('http://localhost:8080/register', {
          email: this.email,
          password: this.password
        });

        if (response.status === 201) {
          alert('Registration successful');
          this.$router.push('/login');
        }
      } catch (error) {
        if (error.response) {
          this.errorMessage = `Registration failed: ${error.response.data}`;
        } else {
          this.errorMessage = 'Registration failed: Unable to connect to the server';
        }
      }
    }
  }
};
</script>

<style scoped>
/* Estilo do container principal */
.background {
  background: url('@/assets/fundo_inicio.jpg') no-repeat center center fixed;
  background-size: cover;
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: 'Rubik', sans-serif;
}

.register-container {
  width: 90%;
  max-width: 500px;
  background: rgba(255, 255, 255, 0.95);
  padding: 30px;
  border-radius: 15px;
  box-shadow: 0 0 30px rgba(0, 0, 0, 0.7);
  text-align: center;
  box-sizing: border-box;
}

/* Estilo do formulário */
form {
  display: flex;
  flex-direction: column;
  align-items: center;
}

form h2 {
  font-size: 2em;
  color: #0044cc;
  margin-bottom: 20px;
}

form input[type="text"],
form input[type="email"],
form input[type="password"] {
  width: 100%;
  padding: 15px;
  margin: 10px 0;
  border: 1px solid #ccc;
  border-radius: 10px;
  box-sizing: border-box;
  font-size: 1em;
}

form button[type="submit"] {
  width: 100%;
  padding: 15px;
  margin: 20px 0;
  background: #0044cc;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 1.2em;
  cursor: pointer;
  transition: background-color 0.3s;
}

form button[type="submit"]:hover {
  background: #003399;
}

/* Estilo das mensagens de erro */
.error-message {
  color: red;
  margin-bottom: 10px;
  display: none; /* Esconde a mensagem por padrão */
}

/* Estilo responsivo */
@media screen and (max-width: 480px) {
  .register-container {
    padding: 20px;
  }
  form input[type="text"],
  form input[type="email"],
  form input[type="password"],
  form button[type="submit"] {
    padding: 10px;
    font-size: 0.9em;
  }
}
</style>
