<template>
  <section id="registerManutencao" class="content-section">
    <h2>REGISTRO DE MANUTENÇÃO</h2>
    <hr>
    <form @submit.prevent="registerManutencao">
      <div class="field">
        <label for="tipo">Tipo de Manutenção</label>
        <select v-model="manutencao.tipo" required>
          <option value="">Selecione o Tipo</option>
          <option value="Troca de óleo">Troca de óleo</option>
          <option value="Reparo de freios">Reparo de freios</option>
          <option value="Troca de pneus">Troca de pneus</option>
          <option value="Revisão geral">Revisão geral</option>
          <option value="Outro">Outro</option>
        </select>
      </div>
      <div class="field">
        <label for="data">Data da Manutenção</label>
        <input type="date" v-model="manutencao.data" required>
      </div>
      <div class="field">
        <label for="valor">Valor da Manutenção</label>
        <input type="number" v-model="manutencao.valor" step="0.01" placeholder="Insira o valor" required>
      </div>
      <div class="field">
        <label for="descricao">Descrição</label>
        <textarea v-model="manutencao.descricao" class="rounded-textarea" placeholder="Insira uma descrição da manutenção" required></textarea>
      </div>
      <!-- <div class="field">
        <label for="placa">Placa da Moto</label>
        <select v-model="manutencao.placa" required>
          <option value="">Selecione a Placa</option>
          <option v-for="moto in this.motos" :key="moto.placa" :value="moto.placa">{{ moto.placa }}</option>
        </select>
      </div>  -->
      <button type="submit">Registrar Manutenção</button>
    </form>
    <div v-if="manutencaoId" class="id-display">
      <p>Manutenção Registrada com Sucesso!</p>
    </div>
  </section>
</template>

<script>
import axios from 'axios';

export default {
  props: ['userId'], // Recebe o userId como prop
  data() {
    return {
      manutencao: {
        tipo: '',
        data: '',
        valor: '',
        descricao: '',
        placa: '' // Novo campo para armazenar a placa selecionada
      },
      manutencaoId: null, // Para armazenar o ID gerado
      motos: [] // Lista de motos disponíveis
    };
  },
  created() {
    this.loadMotosList(); // Carrega a lista de motos ao criar o componente
  },
  methods: {
    // Função para carregar a lista de motos
    loadMotosList() {
      const allMotos = JSON.parse(localStorage.getItem('motos')) || [];
      this.motos = allMotos.filter(moto => moto.userId === this.userId); // Filtra as motos pelo userId

    },
    async registerManutencao() {
      try {
        // Verifica se a placa selecionada existe nas motos registradas
        // const placaExists = this.motos.some(moto => moto.placa === this.manutencao.placa);
        // if (!placaExists) {
        //   alert('Placa selecionada não corresponde a nenhuma moto registrada.');
        //   return;
        // }

        const novaManutencao = {
          tm: this.manutencao.tipo,
          date: this.manutencao.data,
          value: this.manutencao.valor,
          description: this.manutencao.descricao,
          ownerId: this.userId
        };

        const response = await axios.post('http://localhost:8080/maintenance/register', novaManutencao);
        this.manutencaoId = response.data.id;

        this.resetForm();
        alert(`Manutenção registrada com sucesso!`);
      } catch (error) {
        console.error('Error registering maintenance:', error);
        alert('Erro ao registrar manutenção.');
      }
    },
    resetForm() {
      this.manutencao = {
        tipo: '',
        data: '',
        valor: '',
        descricao: '',
        placa: ''
      };
    }
  }
};
</script>

<style scoped>
.content-section {
  width: 100%;
  max-width: 600px;
  text-align: center;
  padding: 20px;
  box-sizing: border-box;
}

.field {
  margin-bottom: 1em;
}

.rounded-textarea {
  width: 100%;
  height: 100px;
  border-radius: 5px;
  padding: 0.5em;
}

button {
  padding: 0.5em 1em;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

.id-display {
  margin-top: 1em;
  padding: 1em;
  background-color: #e0ffe0;
  border: 1px solid #00cc00;
  border-radius: 5px;
}
</style>
