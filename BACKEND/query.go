package main

import (
	"math/rand"
	"strconv"
)

func Querygen() string {

	var genres = []string{
		"acoustic", "afrobeat", "alt-rock", "alternative", "ambient", "anime", "black-metal", "bluegrass", "blues", "bossanova", "brazil", "breakbeat", "british", "cantopop", "chicago-house", "children", "chill", "classical", "club", "comedy", "country", "dance", "dancehall", "death-metal", "deep-house", "detroit-techno", "disco", "disney", "drum-and-bass", "dub", "dubstep", "edm", "electro", "electronic", "emo", "folk", "forro", "french", "funk", "garage", "german", "gospel", "goth", "grindcore", "groove", "grunge", "guitar", "happy", "hard-rock", "hardcore", "hardstyle", "heavy-metal", "hip-hop", "holidays", "honky-tonk", "house", "idm", "indian", "indie", "indie-pop", "industrial", "iranian", "j-dance", "j-idol", "j-pop", "j-rock", "jazz", "k-pop", "kids", "latin", "latino", "malay", "mandopop", "metal", "metal-misc", "metalcore", "minimal-techno", "movies", "mpb", "new-age", "new-release", "opera", "pagode", "party", "philippines-opm", "piano", "pop", "pop-film", "post-dubstep", "power-pop", "progressive-house", "psych-rock", "punk", "punk-rock", "r-n-b", "rainy-day", "reggae", "reggaeton", "road-trip", "rock", "rock-n-roll", "rockabilly", "romance", "sad", "salsa", "samba", "sertanejo", "show-tunes", "singer-songwriter", "ska", "sleep", "songwriter", "soul", "soundtracks", "spanish", "study", "summer", "swedish", "synth-pop", "tango", "techno", "trance", "trip-hop", "turkish", "work-out", "world-music"}

	var char = []string{
		"", "ь", "ч", "ص", "ע", "ر", "ш", "י", "万", "מ", "а", "п", "ж", "十", "h", "t", "e", "四", "ג", "l", "ب", "w", "ظ", "ن", "m", "ت", "צ", "ذ", "百", "k", "خ", "ר", "к", "c", "د", "ו", "ט", "q", "й", " ", "о", "u", "i", "九", "ф", "е", "ط", "و", "ы", "五", "э", "ъ", "ב", "ز", "ц", "ح", "н", "ث", "м", "ש", "م", "щ", "ל", "פ", "س", "ف", "ד", "غ", "б", "ق", "s", "д", "ю", "א", "з", "七", "כ", "я", "j", "هـ", "у", "и", "二", "ח", "z", "л", "ض", "р", "נ", "千", "х", "一", "ش", "ק", "八", "с", "ل", "g", "六", "v", "o", "r", "d", "ع", "ك", "ס", "n", "ז", "ה", "ا", "三", "г", "f", "b", "ي", "ё", "y", "p", "в", "a", "ج", "x", "т", "ת",
	}

	year := rand.Intn(2023-1950) + 1950
	randgen := rand.Intn(len(genres))
	randchar := rand.Intn(len(char))
	query := "genre:" + genres[randgen] + " year:" + strconv.Itoa(year) + " " + char[randchar]

	return query

}
