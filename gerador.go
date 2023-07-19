package main

import(
    "fmt"
    "strings"
    "io/ioutil"
    "os"
    "bufio"
)

func lerTxt(filePath string) []string{
    file, err := os.Open(filePath)
    var slice []string

    if err != nil{
        fmt.Println("Erro linha 16")
    }
    
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        slice = append(slice, line)
    }

    return slice
}

func inserirLinhaEmTxt(message string, filePath string){
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil{
        return
    }
    defer file.Close()

    conteudo, err := ioutil.ReadAll(file)
    if err != nil{
        return
    }
    file.Close()

    file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil{
        return
    }
    defer file.Close()

    newContent := strings.Join([]string{message, string(conteudo)}, "\n")
    _, err = file.WriteString(newContent)
    if err != nil{
        return
    }
}




func gerarCidade(){
    cidades :=  lerTxt("cidades.txt")

    for _, value := range cidades{ 

        str := "INSERT INTO public.cidade(nome) values('"+  value +"');"
        inserirLinhaEmTxt(str,"script.txt")
    }
}

func gerarFuncionario(){
    nomes := lerTxt("nomes.txt")

    for _,value := range nomes{
        str := "INSERT INTO public.funcionarios (nome, cargo, salario)VALUES('"+value+"', '', 1500);"
        inserirLinhaEmTxt(str, "funcDDL.txt")
    }
}

func gerarProduto(){
    for i := 0; i<15; i++{
        str := fmt.Sprintf("INSERT INTO public.produto(descricao, quantidade)VALUES('ProdutoGenerico %d', %d);", i,i)
        inserirLinhaEmTxt(str, "prodDDL.txt")
    }

}

func gerarCliente(){
    nomes := lerTxt("nomes.txt")

    for i,value := range nomes{
        str :=  fmt.Sprintf(" INSERT INTO public.cliente(nome, codcid)VALUES('%s', %d);",value, i)
        inserirLinhaEmTxt(str, "cliDDL.txt")
    }
}

func gerarVenda(){
    for i := 0; i < 15; i++{
        str  := fmt.Sprintf("INSERT INTO public.venda(idprod, idcli, idfunc, valor)VALUES(%d,%d,%d,%d);",i,i,i,i) 
        inserirLinhaEmTxt(str, "vendasDDL.txt")
    }
}

func main(){
    gerarCidade()
    gerarFuncionario()
    gerarProduto()
    gerarCliente()
    gerarVenda()
}
