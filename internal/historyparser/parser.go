package historyparser

func ParseHistory(data *[]byte, bh *BrowseHistory) error {

	var hx = historyXML{}

	if err := hx.parseHistory(data); err != nil {
		return err
	}

	if err := bh.parse(&hx); err != nil {
		return err
	}

	return nil
}
