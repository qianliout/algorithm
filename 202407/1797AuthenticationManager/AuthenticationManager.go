package main

func main() {

}

type AuthenticationManager struct {
	Data map[string]int
	Live int
}

func Constructor(timeToLive int) AuthenticationManager {
	c := AuthenticationManager{
		Data: make(map[string]int),
		Live: timeToLive,
	}

	return c
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.Data[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if this.Data[tokenId] == 0 {
		return
	}
	if this.Data[tokenId]+this.Live <= currentTime {
		return
	}
	this.Data[tokenId] = currentTime
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	cnt := 0
	for _, v := range this.Data {
		if v+this.Live > currentTime {
			cnt++
		}
	}
	return cnt
}
