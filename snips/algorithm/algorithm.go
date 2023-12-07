package algorithm

/*
RemoveExtraSpace 删除字符串中多余的空格
时间复杂度O(N),空间复杂度O(1)
引申一下,可以移除任意字符, .eg __bcc___d_ef__ → bcc_d_ef
*/
func RemoveExtraSpace(sb *[]byte) {
	slow := 0
	fast := 0
	// 去除头部空格
	for (*sb)[fast] == '_' {
		fast++
	}
	// 去除中间部分空格
	for i := fast; i < len(*sb); i++ {
		// 不用从i=1开始, 因为头部的字符已经没了
		// i和i-1比较是精髓
		if i > 1 && (*sb)[i] == (*sb)[i-1] && (*sb)[i] == '_' {
			continue
		} else {
			(*sb)[slow] = (*sb)[i]
			slow++
		}
	}
	// 去除末尾空格
	if slow > 0 && (*sb)[slow-1] == '_' {
		*sb = (*sb)[:slow-1]
		return
	}
	*sb = (*sb)[:slow]
	return
}

func removeExtraSpaceBak(sb *[]byte) {
	slow := 0
	//__bcc___d_ef__ → bcc_d_ef
	for fast := 0; fast < len(*sb); fast++ {
		if (*sb)[fast] == '_' {
			continue
		}
		if slow != 0 {
			(*sb)[slow] = '_'
			slow++
		}

		for fast < len(*sb) && (*sb)[fast] != '_' {
			(*sb)[slow] = (*sb)[fast]
			slow++
			fast++
		}
	}

	*sb = (*sb)[0:slow]
	return
}
