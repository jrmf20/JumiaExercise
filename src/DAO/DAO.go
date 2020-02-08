package DAO

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Customer struct{
	ID int
	Name, Phone string
}

func GetAllCustomers() (customers []Customer, err error){
	db, err := sql.Open("sqlite3", "../sample.db")
	if err != nil {
		return nil, err
	}
	
	rows, err := db.Query("select id,name,phone from customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var phone string
		err = rows.Scan(&id, &name, &phone)
		if err != nil {
			return customers,err
		}
		customers = append(customers,Customer{id,name,phone})
	}
	return customers,nil
}
//
//func inputData(question Question, number int) {
//	
//	var lastid int64
//	//prepara input na bd para questoes
//	insert_Question, err := dbConnection.Prepare("insert into Question(query,answer,created,complete,q_type) values(?,?,CURDATE(),?,?);")
//	if err != nil {log.Fatal(err)}
//	defer insert_Question.Close()
//
//	//prepara input na bd para true statements
//	insert_TrueStatement, err := dbConnection.Prepare("insert into True_Statement(q_id,sentence) values (?, ?);")
//	if err != nil {log.Fatal(err)}
//	defer insert_TrueStatement.Close()
//
//	//prepara input na bd para false statments
//	insert_FalseStatement, err := dbConnection.Prepare("insert into False_Statement(q_id,sentence) values (?, ?);")
//	if err != nil {log.Fatal(err)}
//	defer insert_FalseStatement.Close()
//
//	//prepara input no question topic
//	insert_QT, err := dbConnection.Prepare("INSERT INTO gotest.Question_Topic (q_id, t_id) VALUES(?, 1);")
//	if err != nil {log.Fatal(err)}
//	defer insert_QT.Close()
//	//prepara input para raw misleading answers
//	insert_RMA, err := dbConnection.Prepare("INSERT INTO gotest.Raw_Misleading_Answers (m_answer, q_id) VALUES(?, ?);")
//	if err != nil {log.Fatal(err)}
//	defer insert_RMA.Close()
//
//	//insert q into database
//	var complete byte
//	if question.Answer == "" {
//		complete = 1
//	} else {
//		complete = 0
//	}
//	res, err := insert_Question.Exec(question.Query, question.Answer, complete, question.q_type)
//	if err != nil {log.Fatal(err)}
//	lastid, err = res.LastInsertId()
//	if err != nil {log.Fatal(err)}
//	//insere Question/Topic
//	//TODO: Voltar a olhar para os topicos
//	_, err = insert_QT.Exec(lastid)
//	if err != nil {log.Fatal(err)}
//	switch question.q_type {
//	case 0:
//
//		if question.Answer == "FALSE" {
//			_, err = insert_FalseStatement.Exec(lastid, question.Query)
//			if err != nil {log.Fatal(err)}
//		} else if question.Answer == "TRUE" {
//			_, err = insert_TrueStatement.Exec(lastid, question.Query)
//			if err != nil {log.Fatal(err)}
//		} else {
//			log.Fatal(fmt.Errorf("TF answer options out of bounds. Question Answer: %s", question.Answer))
//		}
//		break
//	case 1:
//		for ma := range question.m_answer {
//			_, err = insert_RMA.Exec(ma, lastid)
//			if err != nil {log.Fatal(err)}
//		}
//		break
//	case 2:
//		//prepara a chamada a função: gotest.INSERT_FGA(:q_id,:t_id,:ma_id,:gap_order,:word_order,:word)
//		insert_FGA, err := dbConnection.Prepare("CALL gotest.INSERT_FGA(?,?,?,?,?,?)")
//		if err != nil {log.Fatal(err)}
//
//		//se a resposta correcta for contextual a pergunta faz dela uma frase afirmativa e guarda-a.
//		if question.Answer_Context {
//			//tentar produzir e inserir true statement correspondente a questão fillgap
////			statement, err := question.ProduceTFStatement()
////			if err != nil {log.Fatal(err)}
//			//insere o resultado da junção anterior
////			_, err = insert_TrueStatement.Exec(lastid, statement)
////			if err != nil {log.Fatal(err)}
//		}
//
//		var gap_answers []string
//		//verifica o numero de espaços na questão
//		n_spaces := strings.Count(question.Query, "____")
//		ma_id := 0
//		//para cada resposta enganadora...
//		for m_answer, context := range question.m_answer {
//			//se a resposta enganadora for contextual
//			if context {
//				//se a questão só tiver um espaço..
//				if n_spaces == 1 {
//					gap_answers = append(gap_answers, m_answer)
//				} else {
//					//TODO: VERIFICADOR AUTOMATICO DE MARCADOR DE INTERVALOS DE RESPOSTAS
//					//Verifica se o nº intervalos corresponde ao nº de respostas dentro da ma
//					if n_separators := strings.Count(m_answer, ","); n_separators != n_spaces-1 {
//						log.Fatal(fmt.Errorf("\nDiferença entre o n de espaços e respostas \n Q: %d \nQuery: %s\nM_answer: %s", number, question.Query, m_answer))
//					} else {
//						gap_answers = strings.Split(m_answer, ",")
//					}
//				}
//				//para cada gap_answer
//				for gap_order, gap_answer := range gap_answers {
//					used_words := strings.Fields(gap_answer)
//					for word_order, word := range used_words {
//						_, err = insert_FGA.Exec(lastid, 1, ma_id, gap_order, word_order, word)
//						if err != nil {log.Fatal(err)}
//					}
//				}
//				//finalmente...
//				//tentar produzir e inserir false statement correspondente a questão fillgap
////				statement, err := question.ProduceTFStatement()
////				if err != nil {log.Fatal(err)}
////				_, err = insert_FalseStatement.Exec(lastid, statement)
////				if err != nil {log.Fatal(err)}
//			}
//
//			//tenta inserir a reposta enganadora
//			_, err = insert_RMA.Exec(m_answer, lastid)
//			if err != nil {log.Fatal(err)}
//			ma_id++
//		}
//		if e:= insert_FGA.Close(); e!= nil{
//			log.Fatal(e)
//		}
//		break
//	case 3:
//		//DO NOTHING
//		break
//	default:
//		log.Fatal(fmt.Errorf("\nQuestion type outside of boundries. how?"))
//	}
//
//	//prepara tag
//	insert_Tag_Desc, err := dbConnection.Prepare("CALL INSERT_TQ(?,?,?)")
//	if err != nil {log.Fatal(err)}
//
////	_, err = insert_Tag_Desc.Exec(lastid, "Title", question.Title)
////	if err != nil {log.Fatal(err)}
////	_, err = insert_Tag_Desc.Exec(lastid, "Learning Objective", question.Learning)
////	if err != nil {log.Fatal(err)}
////	_, err = insert_Tag_Desc.Exec(lastid, "Section Reference", question.Secref)
////	if err != nil {log.Fatal(err)}
////	_, err = insert_Tag_Desc.Exec(lastid, "Difficulty", question.Difficulty)
////	if err != nil {log.Fatal(err)}
////	_, err = insert_Tag_Desc.Exec(lastid, "Bloom's Taxonomy", question.Tax)
////	if err != nil {log.Fatal(err)}
//
//	if e:= insert_Tag_Desc.Close();e!=nil{
//		log.Fatal(e)
//	}
//}
//
////func (r *repository) UpdateQuestion(question *Question){
////	insert_Tag_Desc, err := dbConnection.Prepare("CALL INSERT_TQ(?,?,?)")
////	if err != nil {log.Fatal(err)}
////
////	
////}
