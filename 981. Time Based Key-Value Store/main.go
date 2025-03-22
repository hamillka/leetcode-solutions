type Pair struct {
	val       string
	timeStamp int
}

type TimeMap struct {
	KV map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{
		KV: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.KV[key]; !ok {
		this.KV[key] = make([]Pair, 0)
	}
	this.KV[key] = append(this.KV[key], Pair{
		val:       value,
		timeStamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.KV[key]; !ok {
		return ""
	}
	pairs := this.KV[key]
	if pairs[0].timeStamp > timestamp {
		return ""
	}

	l, r := 0, len(pairs)-1

	for l < r {
		m := (l + r) / 2
		if pairs[m].timeStamp == timestamp {
			return pairs[m].val
		} else if pairs[m].timeStamp > timestamp {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if pairs[l].timeStamp > timestamp {
		return pairs[l-1].val
	}

	return pairs[l].val
}
