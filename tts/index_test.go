package tts

import (
	"os/user"
	"path/filepath"
	"testing"
)

//ALIYUNACCESSID 阿里云 access_id
var ALIYUNACCESSID = ""

//ALIYUNACCESSKEY 阿里云 access_key
var ALIYUNACCESSKEY = ""

func TestGetVoice(t *testing.T) {
	auth := GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
	if _, err := auth.GetVoice("你好,明天见"); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSaveVoice(t *testing.T) {
	myself, _ := user.Current()
	voiceFile := filepath.Join(myself.HomeDir, "Desktop", "a.mp3")
	auth := GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
	auth.TTSConfig.VoiceName = "xiaogang"
	if err := auth.SaveVoice("窗前明月光，地上鞋一双", voiceFile); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetAuth(t *testing.T) {
	auth := GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
	//t.Log(auth.getAuth("你好", "Fri, 20 Oct 2017 01:28:56 GMT"))
	auth.GetVoice("你好")
}

func TestLongText(t *testing.T){
  text:=`
    我家的后面有一个很大的园，相传叫作百草园。现在是早已并屋子一起卖给朱文公的子孙了，连那最末次的相见也已经隔了七八年，其中似乎确凿只有一些野草；但那时却是我的乐园。
　　不必说碧绿的菜畦，光滑的石井栏，高大的皂荚树，紫红的桑葚；也不必说鸣蝉在树叶里长吟，肥胖的黄蜂伏在菜花上，轻捷的叫天子(云雀)忽然从草间直窜向云霄里去了。单是周围的短短的泥墙根一带，就有无限趣味。油蛉在这里低唱，蟋蟀们在这里弹琴。翻开断砖来，有时会遇见蜈蚣；还有斑蝥，倘若用手指按住它的脊梁，便会啪的一声，从后窍喷出一阵烟雾。何首乌藤和木莲藤缠络着，木莲有莲房一般的果实，何首乌有臃肿的根。有人说，何首乌根是有像人形的，吃了便可以成仙，我于是常常拔它起来，牵连不断地拔起来，也曾因此弄坏了泥墙，却从来没有见过有一块根像人样。如果不怕刺，还可以摘到覆盆子，像小珊瑚珠攒成的小球，又酸又甜，色味都比桑葚要好得远。
  `
  myself, _ := user.Current()
  voiceFile := filepath.Join(myself.HomeDir, "Desktop", "a-4.wav")
  auth := GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
  auth.TTSConfig.EncodeType = "wav"
  if err:=auth.SaveLongVoice(text, voiceFile);err!=nil{
    t.Fail()
  }
}
