<template>
  <section id="registroMoto" class="content-section">
    <br> <br> <br>
    <h2>REGISTRO DE MOTO</h2>
    <hr>
    <p>Formulário para registrar sua motocicleta 🏍️</p>
    <form id="motoForm" @submit.prevent="registerMoto">
      <div class="field">
        <label for="marca">Marca da Moto</label>
        <select id="marca" v-model="moto.marca" required @change="filterModelOptions">
          <option value="">Selecione a Marca</option>
          <option v-for="marca in marcas" :key="marca" :value="marca">{{ marca }}</option>
        </select>
      </div>
      <div class="field">
        <label for="modelo">Modelo da Moto</label>
        <select id="modelo" v-model="moto.modelo" required @change="filterCilindradaOptions">
          <option value="">Selecione o Modelo</option>
          <option v-for="modelo in modelos" :key="modelo" :value="modelo">{{ modelo }}</option>
        </select>
      </div>
      <div class="field">
        <label for="ano">Ano de Fabricação</label>
        <input type="number" id="ano" v-model.number="moto.ano" min="1900" max="2024" placeholder="Insira o ano" required>
      </div>
      <div class="field">
        <label for="placa">Placa da Moto</label>
        <input type="text" id="placa" v-model="moto.placa" placeholder="Insira a placa" required>
      </div>
      <div class="field">
        <label for="cor">Cor do Veículo</label>
        <select id="cor" v-model="moto.cor" required>
          <option value="">Selecione a Cor</option>
          <option v-for="cor in cores" :key="cor" :value="cor">{{ cor }}</option>
        </select>
      </div>
      <div class="field">
        <label for="cilindrada">Cilindrada do Veículo</label>
        <select id="cilindrada" v-model="moto.cilindrada" required>
          <option value="">Selecione a Cilindrada</option>
          <option v-for="cilindrada in cilindradas" :key="cilindrada" :value="cilindrada">{{ cilindrada }}</option>
        </select>
      </div>
      <div class="field">
        <label for="kilometragem">Kilometragem Atual do Veículo</label>
        <input type="number" id="kilometragem" v-model.number="moto.kilometragem" placeholder="Insira a kilometragem" required>
      </div>
      <button type="submit">Enviar</button>
    </form>
  </section>
</template>

<script>
export default {
  props: ['userId'],
  data() {
    return {
      moto: {
        marca: '',
        modelo: '',
        ano: '',
        placa: '',
        cor: '',
        cilindrada: '',
        kilometragem: '',
      },
      marcas: ['Honda', 'Yamaha', 'Suzuki', 'BMW', 'Kawasaki'],
      cores: ['Preto', 'Branco', 'Vermelho', 'Azul'],
      modelosPorMarca: {
        'Honda': ['Honda CG', 'Honda Biz', 'Honda PCX', 'Honda CB Twister', 'Honda NXR Bros'],
        'Yamaha': ['Yamaha YBR Factor', 'Yamaha XTZ Crosser', 'Yamaha NMax', 'Yamaha MT', 'Yamaha Fazer'],
        'Suzuki': ['Suzuki Burgman', 'Suzuki V-Strom 650 XT', 'Suzuki GSX-S750', 'Suzuki Intruder', 'Suzuki GSX-R1000'],
        'BMW': ['BMW G 310 R', 'BMW F 750 GS', 'BMW R 1250 GS', 'BMW S 1000 RR'],
        'Kawasaki': ['Kawasaki Ninja', 'Kawasaki Z', 'Kawasaki Versys-X']
      },
      cilindradasPorModelo: {
        'Honda CG': ['125cc', '150cc', '160cc'],
        'Honda Biz': ['125cc'],
        'Honda PCX': ['160cc'],
        'Honda CB Twister': ['250cc', '300cc'],
        'Honda NXR Bros': ['125cc', '150cc', '160cc'],
        'Yamaha YBR Factor': ['125cc', '150cc'],
        'Yamaha XTZ Crosser': ['150cc'],
        'Yamaha NMax': ['160cc'],
        'Yamaha MT': ['300cc', '700cc', '900cc'],
        'Yamaha Fazer': ['150cc', '250cc', '600cc'],
        'Suzuki Burgman': ['125cc'],
        'Suzuki V-Strom 650 XT': ['650cc'],
        'Suzuki GSX-S750': ['750cc'],
        'Suzuki Intruder': ['150cc', '1800cc'],
        'Suzuki GSX-R1000': ['1000cc'],
        'BMW G 310 R': ['310cc'],
        'BMW F 750 GS': ['750cc'],
        'BMW R 1250 GS': ['1250cc'],
        'BMW S 1000 RR': ['1000cc'],
        'Kawasaki Ninja': ['300cc', '400cc', '650cc', '1000cc'],
        'Kawasaki Z': ['300cc', '400cc', '650cc', '1000cc'],
        'Kawasaki Versys-X': ['300cc', '650cc']
      },
      modelos: [],
      cilindradas: []
    };
  },
  methods: {
    filterModelOptions() {
      this.modelos = this.modelosPorMarca[this.moto.marca] || [];
      this.moto.modelo = '';
      this.filterCilindradaOptions();
    },
    filterCilindradaOptions() {
      this.cilindradas = this.cilindradasPorModelo[this.moto.modelo] || [];
      this.moto.cilindrada = '';
    },
    registerMoto() {
      if (!this.userId) {
        alert('User ID is missing');
        return;
      }

      const novaMoto = {
        make: this.moto.marca,
        model: this.moto.modelo,
        year: this.moto.ano,
        plate: this.moto.placa,
        color: this.moto.cor,
        engineCapacity: this.moto.cilindrada,
        mileage: this.moto.kilometragem,
        ownerId: this.userId
      };

      let motos = JSON.parse(localStorage.getItem('motos')) || [];
      motos.push(novaMoto);
      localStorage.setItem('motos', JSON.stringify(motos));

      alert('Moto registrada com sucesso!');
      this.resetForm();
    },
    resetForm() {
      this.moto = {
        marca: '',
        modelo: '',
        ano: '',
        placa: '',
        cor: '',
        cilindrada: '',
        kilometragem: '',
      };
      this.modelos = [];
      this.cilindradas = [];
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

#registerMoto {
  display: flex;
  flex-direction: column;
  align-items: center;
}

#motoForm {
  width: 100%;
  height: 70%;
  max-width: 500px;
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.field {
  margin-bottom: 15px;
  text-align: left;
}

.field label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

.field input, .field select {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 5px;
}

#motoForm button {
  width: 100%;
  padding: 12px;
  background-color: #0044cc;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 1em;
  font-weight: bold;
  cursor: pointer;
  transition: all 0.3s ease;
}

#motoForm button:hover {
  background-color: white;
  color: #0044cc;
  border: 2px solid #0044cc;
  transform: scale(1.05);
}
</style>
