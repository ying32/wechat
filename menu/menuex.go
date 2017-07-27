// menuex.go by:ying32
package menu

type TWXMenu struct {
	menu *Menu
	// 微信接口
	// 顶级菜单
	topItems []*TMenuItem
}

// NewWXMenu 新建菜单
func NewWXMenu(menu *Menu) *TWXMenu {
	m := new(TWXMenu)
	m.menu = menu
	return m
}

// Add 添加一个顶级菜单
func (m *TWXMenu) Add(name string) *TMenuItem {
	item := new(TMenuItem)
	item.name = name
	item.item = new(Button)
	m.topItems = append(m.topItems, item)
	return item
}

func (m *TWXMenu) GetMenuButtons() []*Button {
	result := make([]*Button, len(m.topItems))
	for i := 0; i < len(result); i++ {
		m.topItems[i].SetSubMenu()
		result[i] = m.topItems[i].GetItem()
	}
	return result
}

// Submit 提交菜单
func (m *TWXMenu) Submit() error {
	return m.menu.SetMenu(m.GetMenuButtons())
}

// TMenuItem 顶级菜单项
type TMenuItem struct {
	name     string
	item     *Button
	subItems []*Button
}

func (s *TMenuItem) SetSubMenu() *TMenuItem {
	s.item.SetSubButton(s.name, s.subItems)
	return s
}

func (s *TMenuItem) GetItem() *Button {
	return s.item
}

func (s *TMenuItem) Add() *Button {
	btn := new(Button)
	s.subItems = append(s.subItems, btn)
	return btn
}

func (s *TMenuItem) AddClick(name, value string) *Button {
	btn := s.Add()
	btn.SetClickButton(name, value)
	return btn
}

func (s *TMenuItem) AddView(name, value string) *Button {
	btn := s.Add()
	btn.SetViewButton(name, value)
	return btn
}

func (s *TMenuItem) AddScanCodePush(name, value string) *Button {
	btn := s.Add()
	btn.SetScanCodePushButton(name, value)
	return btn
}

func (s *TMenuItem) AddScanCodeWaitMsg(name, value string) *Button {
	btn := s.Add()
	btn.SetScanCodeWaitMsgButton(name, value)
	return btn
}

func (s *TMenuItem) AddPicSysPhoto(name, value string) *Button {
	btn := s.Add()
	btn.SetPicSysPhotoButton(name, value)
	return btn
}

func (s *TMenuItem) AddPicPhotoOrAlbum(name, value string) *Button {
	btn := s.Add()
	btn.SetPicPhotoOrAlbumButton(name, value)
	return btn
}

func (s *TMenuItem) AddPicWeixin(name, value string) *Button {
	btn := s.Add()
	btn.SetPicWeixinButton(name, value)
	return btn
}

func (s *TMenuItem) AddLocationSelect(name, value string) *Button {
	btn := s.Add()
	btn.SetLocationSelectButton(name, value)
	return btn
}

func (s *TMenuItem) AddMediaID(name, value string) *Button {
	btn := s.Add()
	btn.SetMediaIDButton(name, value)
	return btn
}

func (s *TMenuItem) AddViewLimited(name, value string) *Button {
	btn := s.Add()
	btn.SetViewLimitedButton(name, value)
	return btn
}
