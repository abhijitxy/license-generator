

<h1 align="center">Generador de Licencias</h1>


<p align="center">
  <img src="https://drive.google.com/uc?export=view&id=1IT2o8A9gqdbTmH4uxqBYs5pp2-qJ0jdS" alt="Demostración del Generador de Licencias" style="max-width:100%;"/>
</p>

### Descripción General

Con el Generador de Licencias 🚀, puedes optimizar tu flujo de trabajo sin esfuerzo. En lugar de navegar por la interfaz de GitHub 🖥️, simplemente ejecutas un comando en tu terminal para generar y agregar la licencia que desees directamente en tu proyecto 🛠️. 

Esta herramienta elimina la necesidad de pasos manuales tediosos, permitiéndote concentrarte más en escribir código y menos en tareas administrativas 🧑‍💻. Ya sea que estés configurando un repositorio nuevo o actualizando uno existente, el Generador de Licencias hace que el proceso sea rápido y sin complicaciones 🎉.

### Instalación

Descarga el binario más reciente desde la [página de lanzamientos](https://github.com/abhijitxy/license-generator/releases).

```bash
# Para Linux
wget https://github.com/abhijitxy/license-generator/releases/download/v1.0.0/license-generator-linux
chmod +x license-generator-linux
mv license-generator-linux /usr/local/bin/license-gen
```

### Uso

```bash
license-gen
```

### Contribución

- Crea un fork del repositorio
- Crea una rama
- Instala las dependencias
```bash
go mod tidy
```
- Compila
```bash
 go build -o license-generator
 ```
- Confirma tus cambios y envíalos a tu rama
```bash
git commit -m "made an awesomeFix"
git push origin fix/awesomeFix
```
- Abre una pull request
