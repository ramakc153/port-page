async function getClientIP(){
    try{
        const response = await fetch('https://api.ipify.org?format=json')
        if (!response.ok){
            throw new Error(`HTTP Error status: ${response.status}`)
        }
        const result = await response.json()
        // console.log(result.ip)
        return result.ip
    } catch(error){
        console.error(error)
        return null
    }
}

async function getLocation() {
    const clientIp = await getClientIP()
    try{
        const response = await fetch(`http://ip-api.com/json/${clientIp}`)
        if (!response.ok){
            throw new Error(`HTTP Error status: ${response.status}`)
        }
        const result = await response.json()
        // console.log(result)
        return result
    } catch(error){
        console.error(error)
        return null
    }
}

getLocation()

async function getWeather(){
    const clientIp = await getLocation()
    try{
        const lat = clientIp.lat
        const lon = clientIp.lon
        const response = await fetch(`https://api.open-meteo.com/v1/forecast?latitude=${lat}&longitude=${lon}&current=temperature_2m`)
        if (!response.ok){
            throw new Error(`HTTP Error status: ${response.status}`)
        }
        const result = await response.json()
        // console.log(result.current.temperature_2m)
        let temperatureHtml = document.createElement("p")
        temperatureHtml.innerHTML = `your current temperature is: ${result.current.temperature_2m}°C`
        document.body.appendChild(temperatureHtml)
        return result
    } catch(error){
        console.error(error)
        return null
    }
}

getWeather()

