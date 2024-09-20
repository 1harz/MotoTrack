<template>
  <section id="listMotos" class="content-section">
    <h2>LISTA DE MOTOS</h2>
    <hr>
    <button @click="clearMotosList">Limpar Lista de Motos</button> <br> <br>
    <div v-if="motos.length === 0">
      <p>Nenhuma moto disponível.</p>
    </div>
    <div v-for="(moto, index) in motos" :key="index" class="moto-card">
      <h3>{{ moto.make }} - {{ moto.model }}</h3>
      <p><strong>Ano:</strong> {{ moto.year }}</p>
      <p><strong>Placa:</strong> {{ moto.plate }}</p>
      <p><strong>Cor:</strong> {{ moto.color }}</p>
      <p><strong>Cilindrada:</strong> {{ moto.engineCapacity }}</p>
      <p><strong>Kilometragem:</strong> {{ moto.mileage }} km</p>
    </div>
  </section>
</template>

<script>
import axios from 'axios';

export default {
  props: ['userId'], // Recebe o userId como prop
  data() {
    return {
      motos: [] // Lista de motos
    };
  },
  created() {
    this.loadMotosList(); // Carrega a lista de motos ao criar o componente
  },
  methods: {
    // Função para carregar a lista de motos
    async loadMotosList() {
      try {
        const response = await axios.get(`http://localhost:8080/motorcicles?userId=${this.userId}`);
        this.motos = response.data;
      } catch (error) {
        console.error('Error loading motos:', error);
        alert('Erro ao carregar a lista de motos.');
      }
    },
    // Função para limpar a lista de motos
    clearMotosList() {
      // Enviar uma requisição DELETE para limpar as motos do usuário
      axios.delete(`http://localhost:8080/motorcicles?userId=${this.userId}`)
        .then(() => {
          this.motos = [];
        })
        .catch(error => {
          console.error('Error clearing motos:', error);
          alert('Erro ao limpar a lista de motos.');
        });
    }
  }
};
</script>

<style scoped>
.content-section {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.moto-card {
  margin-bottom: 20px;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #fff;
}

.moto-card h3 {
  margin: 0 0 10px 0;
}

button {
  padding: 10px;
  margin-top: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
