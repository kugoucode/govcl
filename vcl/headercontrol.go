
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type THeaderControl struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewHeaderControl(owner IComponent) *THeaderControl {
    h := new(THeaderControl)
    h.instance = HeaderControl_Create(CheckPtr(owner))
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsHeaderControl(obj interface{}) *THeaderControl {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &THeaderControl{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsHeaderControl.
func HeaderControlFromInst(inst uintptr) *THeaderControl {
    return AsHeaderControl(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsHeaderControl.
func HeaderControlFromObj(obj IObject) *THeaderControl {
    return AsHeaderControl(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsHeaderControl.
func HeaderControlFromUnsafePointer(ptr unsafe.Pointer) *THeaderControl {
    return AsHeaderControl(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (h *THeaderControl) Free() {
    if h.instance != 0 {
        HeaderControl_Free(h.instance)
        h.instance, h.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (h *THeaderControl) Instance() uintptr {
    return h.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (h *THeaderControl) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (h *THeaderControl) IsValid() bool {
    return h.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (h *THeaderControl) Is() TIs {
    return TIs(h.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (h *THeaderControl) As() TAs {
//    return TAs(h.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func THeaderControlClass() TClass {
    return HeaderControl_StaticClassType()
}

func (h *THeaderControl) FlipChildren(AllLevels bool) {
    HeaderControl_FlipChildren(h.instance, AllLevels)
}

// 是否可以获得焦点。
func (h *THeaderControl) CanFocus() bool {
    return HeaderControl_CanFocus(h.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (h *THeaderControl) ContainsControl(Control IControl) bool {
    return HeaderControl_ContainsControl(h.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (h *THeaderControl) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(HeaderControl_ControlAtPos(h.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (h *THeaderControl) DisableAlign() {
    HeaderControl_DisableAlign(h.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (h *THeaderControl) EnableAlign() {
    HeaderControl_EnableAlign(h.instance)
}

// 查找子控件。
//
// Find sub controls.
func (h *THeaderControl) FindChildControl(ControlName string) *TControl {
    return AsControl(HeaderControl_FindChildControl(h.instance, ControlName))
}

// 返回是否获取焦点。
//
// Return to get focus.
func (h *THeaderControl) Focused() bool {
    return HeaderControl_Focused(h.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (h *THeaderControl) HandleAllocated() bool {
    return HeaderControl_HandleAllocated(h.instance)
}

// 插入一个控件。
//
// Insert a control.
func (h *THeaderControl) InsertControl(AControl IControl) {
    HeaderControl_InsertControl(h.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (h *THeaderControl) Invalidate() {
    HeaderControl_Invalidate(h.instance)
}

// 移除一个控件。
//
// Remove a control.
func (h *THeaderControl) RemoveControl(AControl IControl) {
    HeaderControl_RemoveControl(h.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (h *THeaderControl) Realign() {
    HeaderControl_Realign(h.instance)
}

// 重绘。
//
// Repaint.
func (h *THeaderControl) Repaint() {
    HeaderControl_Repaint(h.instance)
}

// 按比例缩放。
//
// Scale by.
func (h *THeaderControl) ScaleBy(M int32, D int32) {
    HeaderControl_ScaleBy(h.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (h *THeaderControl) ScrollBy(DeltaX int32, DeltaY int32) {
    HeaderControl_ScrollBy(h.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (h *THeaderControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HeaderControl_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (h *THeaderControl) SetFocus() {
    HeaderControl_SetFocus(h.instance)
}

// 控件更新。
//
// Update.
func (h *THeaderControl) Update() {
    HeaderControl_Update(h.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (h *THeaderControl) BringToFront() {
    HeaderControl_BringToFront(h.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (h *THeaderControl) ClientToScreen(Point TPoint) TPoint {
    return HeaderControl_ClientToScreen(h.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (h *THeaderControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ClientToParent(h.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (h *THeaderControl) Dragging() bool {
    return HeaderControl_Dragging(h.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (h *THeaderControl) HasParent() bool {
    return HeaderControl_HasParent(h.instance)
}

// 隐藏控件。
//
// Hidden control.
func (h *THeaderControl) Hide() {
    HeaderControl_Hide(h.instance)
}

// 发送一个消息。
//
// Send a message.
func (h *THeaderControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HeaderControl_Perform(h.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (h *THeaderControl) Refresh() {
    HeaderControl_Refresh(h.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (h *THeaderControl) ScreenToClient(Point TPoint) TPoint {
    return HeaderControl_ScreenToClient(h.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (h *THeaderControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ParentToClient(h.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (h *THeaderControl) SendToBack() {
    HeaderControl_SendToBack(h.instance)
}

// 显示控件。
//
// Show control.
func (h *THeaderControl) Show() {
    HeaderControl_Show(h.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (h *THeaderControl) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return HeaderControl_GetTextBuf(h.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (h *THeaderControl) GetTextLen() int32 {
    return HeaderControl_GetTextLen(h.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (h *THeaderControl) SetTextBuf(Buffer string) {
    HeaderControl_SetTextBuf(h.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (h *THeaderControl) FindComponent(AName string) *TComponent {
    return AsComponent(HeaderControl_FindComponent(h.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (h *THeaderControl) GetNamePath() string {
    return HeaderControl_GetNamePath(h.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (h *THeaderControl) Assign(Source IObject) {
    HeaderControl_Assign(h.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (h *THeaderControl) ClassType() TClass {
    return HeaderControl_ClassType(h.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (h *THeaderControl) ClassName() string {
    return HeaderControl_ClassName(h.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (h *THeaderControl) InstanceSize() int32 {
    return HeaderControl_InstanceSize(h.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (h *THeaderControl) InheritsFrom(AClass TClass) bool {
    return HeaderControl_InheritsFrom(h.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (h *THeaderControl) Equals(Obj IObject) bool {
    return HeaderControl_Equals(h.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (h *THeaderControl) GetHashCode() int32 {
    return HeaderControl_GetHashCode(h.instance)
}

// 文本类信息。
//
// Text information.
func (h *THeaderControl) ToString() string {
    return HeaderControl_ToString(h.instance)
}

func (h *THeaderControl) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    HeaderControl_AnchorToNeighbour(h.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (h *THeaderControl) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    HeaderControl_AnchorParallel(h.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (h *THeaderControl) AnchorHorizontalCenterTo(ASibling IControl) {
    HeaderControl_AnchorHorizontalCenterTo(h.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (h *THeaderControl) AnchorVerticalCenterTo(ASibling IControl) {
    HeaderControl_AnchorVerticalCenterTo(h.instance, CheckPtr(ASibling))
}

func (h *THeaderControl) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    HeaderControl_AnchorAsAlign(h.instance, ATheAlign , ASpace)
}

func (h *THeaderControl) AnchorClient(ASpace int32) {
    HeaderControl_AnchorClient(h.instance, ASpace)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (h *THeaderControl) Align() TAlign {
    return HeaderControl_GetAlign(h.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (h *THeaderControl) SetAlign(value TAlign) {
    HeaderControl_SetAlign(h.instance, value)
}

// 获取四个角位置的锚点。
func (h *THeaderControl) Anchors() TAnchors {
    return HeaderControl_GetAnchors(h.instance)
}

// 设置四个角位置的锚点。
func (h *THeaderControl) SetAnchors(value TAnchors) {
    HeaderControl_SetAnchors(h.instance, value)
}

// 获取边框的宽度。
func (h *THeaderControl) BorderWidth() int32 {
    return HeaderControl_GetBorderWidth(h.instance)
}

// 设置边框的宽度。
func (h *THeaderControl) SetBorderWidth(value int32) {
    HeaderControl_SetBorderWidth(h.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (h *THeaderControl) DoubleBuffered() bool {
    return HeaderControl_GetDoubleBuffered(h.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (h *THeaderControl) SetDoubleBuffered(value bool) {
    HeaderControl_SetDoubleBuffered(h.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (h *THeaderControl) DragCursor() TCursor {
    return HeaderControl_GetDragCursor(h.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (h *THeaderControl) SetDragCursor(value TCursor) {
    HeaderControl_SetDragCursor(h.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (h *THeaderControl) DragKind() TDragKind {
    return HeaderControl_GetDragKind(h.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (h *THeaderControl) SetDragKind(value TDragKind) {
    HeaderControl_SetDragKind(h.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (h *THeaderControl) DragMode() TDragMode {
    return HeaderControl_GetDragMode(h.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (h *THeaderControl) SetDragMode(value TDragMode) {
    HeaderControl_SetDragMode(h.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (h *THeaderControl) Enabled() bool {
    return HeaderControl_GetEnabled(h.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (h *THeaderControl) SetEnabled(value bool) {
    HeaderControl_SetEnabled(h.instance, value)
}

// 获取字体。
//
// Get Font.
func (h *THeaderControl) Font() *TFont {
    return AsFont(HeaderControl_GetFont(h.instance))
}

// 设置字体。
//
// Set Font.
func (h *THeaderControl) SetFont(value *TFont) {
    HeaderControl_SetFont(h.instance, CheckPtr(value))
}

// 获取图标索引列表对象。
func (h *THeaderControl) Images() *TImageList {
    return AsImageList(HeaderControl_GetImages(h.instance))
}

// 设置图标索引列表对象。
func (h *THeaderControl) SetImages(value IComponent) {
    HeaderControl_SetImages(h.instance, CheckPtr(value))
}

// 获取约束控件大小。
func (h *THeaderControl) Constraints() *TSizeConstraints {
    return AsSizeConstraints(HeaderControl_GetConstraints(h.instance))
}

// 设置约束控件大小。
func (h *THeaderControl) SetConstraints(value *TSizeConstraints) {
    HeaderControl_SetConstraints(h.instance, CheckPtr(value))
}

func (h *THeaderControl) Sections() *THeaderSections {
    return AsHeaderSections(HeaderControl_GetSections(h.instance))
}

func (h *THeaderControl) SetSections(value *THeaderSections) {
    HeaderControl_SetSections(h.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (h *THeaderControl) ShowHint() bool {
    return HeaderControl_GetShowHint(h.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (h *THeaderControl) SetShowHint(value bool) {
    HeaderControl_SetShowHint(h.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (h *THeaderControl) ParentDoubleBuffered() bool {
    return HeaderControl_GetParentDoubleBuffered(h.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (h *THeaderControl) SetParentDoubleBuffered(value bool) {
    HeaderControl_SetParentDoubleBuffered(h.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (h *THeaderControl) ParentFont() bool {
    return HeaderControl_GetParentFont(h.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (h *THeaderControl) SetParentFont(value bool) {
    HeaderControl_SetParentFont(h.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (h *THeaderControl) ParentShowHint() bool {
    return HeaderControl_GetParentShowHint(h.instance)
}

// 设置以父容器的ShowHint属性为准。
func (h *THeaderControl) SetParentShowHint(value bool) {
    HeaderControl_SetParentShowHint(h.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (h *THeaderControl) PopupMenu() *TPopupMenu {
    return AsPopupMenu(HeaderControl_GetPopupMenu(h.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (h *THeaderControl) SetPopupMenu(value IComponent) {
    HeaderControl_SetPopupMenu(h.instance, CheckPtr(value))
}

// 获取控件可视。
//
// Get the control visible.
func (h *THeaderControl) Visible() bool {
    return HeaderControl_GetVisible(h.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (h *THeaderControl) SetVisible(value bool) {
    HeaderControl_SetVisible(h.instance, value)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (h *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
    HeaderControl_SetOnContextPopup(h.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (h *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
    HeaderControl_SetOnDragDrop(h.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (h *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
    HeaderControl_SetOnDragOver(h.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (h *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
    HeaderControl_SetOnEndDock(h.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (h *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
    HeaderControl_SetOnEndDrag(h.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (h *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
    HeaderControl_SetOnMouseDown(h.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (h *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
    HeaderControl_SetOnMouseEnter(h.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (h *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
    HeaderControl_SetOnMouseLeave(h.instance, fn)
}

// 设置鼠标移动事件。
func (h *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
    HeaderControl_SetOnMouseMove(h.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (h *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
    HeaderControl_SetOnMouseUp(h.instance, fn)
}

// 设置大小被改变事件。
func (h *THeaderControl) SetOnResize(fn TNotifyEvent) {
    HeaderControl_SetOnResize(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionClick(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionClick(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionResize(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionResize(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionTrack(fn TSectionTrackEvent) {
    HeaderControl_SetOnSectionTrack(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
    HeaderControl_SetOnSectionDrag(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
    HeaderControl_SetOnSectionEndDrag(h.instance, fn)
}

// 获取画布。
func (h *THeaderControl) Canvas() *TCanvas {
    return AsCanvas(HeaderControl_GetCanvas(h.instance))
}

// 获取依靠客户端总数。
func (h *THeaderControl) DockClientCount() int32 {
    return HeaderControl_GetDockClientCount(h.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (h *THeaderControl) DockSite() bool {
    return HeaderControl_GetDockSite(h.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (h *THeaderControl) SetDockSite(value bool) {
    HeaderControl_SetDockSite(h.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (h *THeaderControl) MouseInClient() bool {
    return HeaderControl_GetMouseInClient(h.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (h *THeaderControl) VisibleDockClientCount() int32 {
    return HeaderControl_GetVisibleDockClientCount(h.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (h *THeaderControl) Brush() *TBrush {
    return AsBrush(HeaderControl_GetBrush(h.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (h *THeaderControl) ControlCount() int32 {
    return HeaderControl_GetControlCount(h.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (h *THeaderControl) Handle() HWND {
    return HeaderControl_GetHandle(h.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (h *THeaderControl) ParentWindow() HWND {
    return HeaderControl_GetParentWindow(h.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (h *THeaderControl) SetParentWindow(value HWND) {
    HeaderControl_SetParentWindow(h.instance, value)
}

func (h *THeaderControl) Showing() bool {
    return HeaderControl_GetShowing(h.instance)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (h *THeaderControl) TabOrder() TTabOrder {
    return HeaderControl_GetTabOrder(h.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (h *THeaderControl) SetTabOrder(value TTabOrder) {
    HeaderControl_SetTabOrder(h.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (h *THeaderControl) TabStop() bool {
    return HeaderControl_GetTabStop(h.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (h *THeaderControl) SetTabStop(value bool) {
    HeaderControl_SetTabStop(h.instance, value)
}

// 获取使用停靠管理。
func (h *THeaderControl) UseDockManager() bool {
    return HeaderControl_GetUseDockManager(h.instance)
}

// 设置使用停靠管理。
func (h *THeaderControl) SetUseDockManager(value bool) {
    HeaderControl_SetUseDockManager(h.instance, value)
}

func (h *THeaderControl) Action() *TAction {
    return AsAction(HeaderControl_GetAction(h.instance))
}

func (h *THeaderControl) SetAction(value IComponent) {
    HeaderControl_SetAction(h.instance, CheckPtr(value))
}

func (h *THeaderControl) BoundsRect() TRect {
    return HeaderControl_GetBoundsRect(h.instance)
}

func (h *THeaderControl) SetBoundsRect(value TRect) {
    HeaderControl_SetBoundsRect(h.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (h *THeaderControl) ClientHeight() int32 {
    return HeaderControl_GetClientHeight(h.instance)
}

// 设置客户区高度。
//
// Set client height.
func (h *THeaderControl) SetClientHeight(value int32) {
    HeaderControl_SetClientHeight(h.instance, value)
}

func (h *THeaderControl) ClientOrigin() TPoint {
    return HeaderControl_GetClientOrigin(h.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (h *THeaderControl) ClientRect() TRect {
    return HeaderControl_GetClientRect(h.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (h *THeaderControl) ClientWidth() int32 {
    return HeaderControl_GetClientWidth(h.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (h *THeaderControl) SetClientWidth(value int32) {
    HeaderControl_SetClientWidth(h.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (h *THeaderControl) ControlState() TControlState {
    return HeaderControl_GetControlState(h.instance)
}

// 设置控件状态。
//
// Set control state.
func (h *THeaderControl) SetControlState(value TControlState) {
    HeaderControl_SetControlState(h.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (h *THeaderControl) ControlStyle() TControlStyle {
    return HeaderControl_GetControlStyle(h.instance)
}

// 设置控件样式。
//
// Set control style.
func (h *THeaderControl) SetControlStyle(value TControlStyle) {
    HeaderControl_SetControlStyle(h.instance, value)
}

func (h *THeaderControl) Floating() bool {
    return HeaderControl_GetFloating(h.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (h *THeaderControl) Parent() *TWinControl {
    return AsWinControl(HeaderControl_GetParent(h.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (h *THeaderControl) SetParent(value IWinControl) {
    HeaderControl_SetParent(h.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (h *THeaderControl) Left() int32 {
    return HeaderControl_GetLeft(h.instance)
}

// 设置左边位置。
//
// Set Left position.
func (h *THeaderControl) SetLeft(value int32) {
    HeaderControl_SetLeft(h.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (h *THeaderControl) Top() int32 {
    return HeaderControl_GetTop(h.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (h *THeaderControl) SetTop(value int32) {
    HeaderControl_SetTop(h.instance, value)
}

// 获取宽度。
//
// Get width.
func (h *THeaderControl) Width() int32 {
    return HeaderControl_GetWidth(h.instance)
}

// 设置宽度。
//
// Set width.
func (h *THeaderControl) SetWidth(value int32) {
    HeaderControl_SetWidth(h.instance, value)
}

// 获取高度。
//
// Get height.
func (h *THeaderControl) Height() int32 {
    return HeaderControl_GetHeight(h.instance)
}

// 设置高度。
//
// Set height.
func (h *THeaderControl) SetHeight(value int32) {
    HeaderControl_SetHeight(h.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (h *THeaderControl) Cursor() TCursor {
    return HeaderControl_GetCursor(h.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (h *THeaderControl) SetCursor(value TCursor) {
    HeaderControl_SetCursor(h.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (h *THeaderControl) Hint() string {
    return HeaderControl_GetHint(h.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (h *THeaderControl) SetHint(value string) {
    HeaderControl_SetHint(h.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (h *THeaderControl) ComponentCount() int32 {
    return HeaderControl_GetComponentCount(h.instance)
}

// 获取组件索引。
//
// Get component index.
func (h *THeaderControl) ComponentIndex() int32 {
    return HeaderControl_GetComponentIndex(h.instance)
}

// 设置组件索引。
//
// Set component index.
func (h *THeaderControl) SetComponentIndex(value int32) {
    HeaderControl_SetComponentIndex(h.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (h *THeaderControl) Owner() *TComponent {
    return AsComponent(HeaderControl_GetOwner(h.instance))
}

// 获取组件名称。
//
// Get the component name.
func (h *THeaderControl) Name() string {
    return HeaderControl_GetName(h.instance)
}

// 设置组件名称。
//
// Set the component name.
func (h *THeaderControl) SetName(value string) {
    HeaderControl_SetName(h.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (h *THeaderControl) Tag() int {
    return HeaderControl_GetTag(h.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (h *THeaderControl) SetTag(value int) {
    HeaderControl_SetTag(h.instance, value)
}

// 获取左边锚点。
func (h *THeaderControl) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(HeaderControl_GetAnchorSideLeft(h.instance))
}

// 设置左边锚点。
func (h *THeaderControl) SetAnchorSideLeft(value *TAnchorSide) {
    HeaderControl_SetAnchorSideLeft(h.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (h *THeaderControl) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(HeaderControl_GetAnchorSideTop(h.instance))
}

// 设置顶边锚点。
func (h *THeaderControl) SetAnchorSideTop(value *TAnchorSide) {
    HeaderControl_SetAnchorSideTop(h.instance, CheckPtr(value))
}

// 获取右边锚点。
func (h *THeaderControl) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(HeaderControl_GetAnchorSideRight(h.instance))
}

// 设置右边锚点。
func (h *THeaderControl) SetAnchorSideRight(value *TAnchorSide) {
    HeaderControl_SetAnchorSideRight(h.instance, CheckPtr(value))
}

// 获取底边锚点。
func (h *THeaderControl) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(HeaderControl_GetAnchorSideBottom(h.instance))
}

// 设置底边锚点。
func (h *THeaderControl) SetAnchorSideBottom(value *TAnchorSide) {
    HeaderControl_SetAnchorSideBottom(h.instance, CheckPtr(value))
}

func (h *THeaderControl) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(HeaderControl_GetChildSizing(h.instance))
}

func (h *THeaderControl) SetChildSizing(value *TControlChildSizing) {
    HeaderControl_SetChildSizing(h.instance, CheckPtr(value))
}

// 获取边框间距。
func (h *THeaderControl) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(HeaderControl_GetBorderSpacing(h.instance))
}

// 设置边框间距。
func (h *THeaderControl) SetBorderSpacing(value *TControlBorderSpacing) {
    HeaderControl_SetBorderSpacing(h.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (h *THeaderControl) DockClients(Index int32) *TControl {
    return AsControl(HeaderControl_GetDockClients(h.instance, Index))
}

// 获取指定索引子控件。
func (h *THeaderControl) Controls(Index int32) *TControl {
    return AsControl(HeaderControl_GetControls(h.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (h *THeaderControl) Components(AIndex int32) *TComponent {
    return AsComponent(HeaderControl_GetComponents(h.instance, AIndex))
}

// 获取锚侧面。
func (h *THeaderControl) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(HeaderControl_GetAnchorSide(h.instance, AKind))
}

