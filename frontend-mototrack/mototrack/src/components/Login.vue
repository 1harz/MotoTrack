<template>
  <div class="fundo">
    <div class="bg-text">
      <h1 class="Welcome">
        Welcome to 
        <div class="Welcome_Name">
          <a target="_blank" href="info.html">MotoTrack</a>
        </div>
      </h1>
      <div class="form-container">
        <form @submit.prevent="signup">
          <div class="form-group">
            <input type="email" v-model="email" class="form-control" placeholder="Email" required>
          </div>
          <div class="form-group">
            <input type="password" v-model="password" class="form-control" placeholder="Password" required>
          </div>
          <button type="submit" class="btn btn-dark btn-block">Sign in</button> <br> <br>
          <div>Dont have an account? <router-link to="/register">Sign up Here</router-link></div>
        </form>
        <hr style="color: white;">
      </div>
      <p>By clicking continue, you agree to our <a target="_blank" href="terms.html">Terms of Service and Privacy Policy</a></p>
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
    };
  },
  mounted() {
    const userId = localStorage.getItem('userId');
    if (userId) {
      this.$router.push('/control'); // Redireciona para /control se já estiver logado
    }
  },
  methods: {
    async signup() {
      try {
        const response = await axios.post('http://localhost:8080/login', {
          email: this.email,
          password: this.password
        });

        if (response.status === 200) {
          alert('Login successful');
          localStorage.setItem('userId', response.data.id); // Salvar userId no localStorage
          this.$router.push('/control');
        }
      } catch (error) {
        alert('Invalid email or password');
      }
    }
  }
};
</script>

<style>
@import '../assets/inicio.css';
</style>
