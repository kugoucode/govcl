
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

type TShape struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewShape(owner IComponent) *TShape {
    s := new(TShape)
    s.instance = Shape_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsShape(obj interface{}) *TShape {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TShape{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsShape.
func ShapeFromInst(inst uintptr) *TShape {
    return AsShape(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsShape.
func ShapeFromObj(obj IObject) *TShape {
    return AsShape(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsShape.
func ShapeFromUnsafePointer(ptr unsafe.Pointer) *TShape {
    return AsShape(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (s *TShape) Free() {
    if s.instance != 0 {
        Shape_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (s *TShape) Instance() uintptr {
    return s.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (s *TShape) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (s *TShape) IsValid() bool {
    return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (s *TShape) Is() TIs {
    return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (s *TShape) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TShapeClass() TClass {
    return Shape_StaticClassType()
}

// 将控件置于最前。
//
// Bring the control to the front.
func (s *TShape) BringToFront() {
    Shape_BringToFront(s.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (s *TShape) ClientToScreen(Point TPoint) TPoint {
    return Shape_ClientToScreen(s.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (s *TShape) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Shape_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (s *TShape) Dragging() bool {
    return Shape_Dragging(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TShape) HasParent() bool {
    return Shape_HasParent(s.instance)
}

// 隐藏控件。
//
// Hidden control.
func (s *TShape) Hide() {
    Shape_Hide(s.instance)
}

// 要求重绘。
//
// Redraw.
func (s *TShape) Invalidate() {
    Shape_Invalidate(s.instance)
}

// 发送一个消息。
//
// Send a message.
func (s *TShape) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Shape_Perform(s.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (s *TShape) Refresh() {
    Shape_Refresh(s.instance)
}

// 重绘。
//
// Repaint.
func (s *TShape) Repaint() {
    Shape_Repaint(s.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (s *TShape) ScreenToClient(Point TPoint) TPoint {
    return Shape_ScreenToClient(s.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (s *TShape) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Shape_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (s *TShape) SendToBack() {
    Shape_SendToBack(s.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (s *TShape) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Shape_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// 显示控件。
//
// Show control.
func (s *TShape) Show() {
    Shape_Show(s.instance)
}

// 控件更新。
//
// Update.
func (s *TShape) Update() {
    Shape_Update(s.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (s *TShape) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Shape_GetTextBuf(s.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (s *TShape) GetTextLen() int32 {
    return Shape_GetTextLen(s.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (s *TShape) SetTextBuf(Buffer string) {
    Shape_SetTextBuf(s.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TShape) FindComponent(AName string) *TComponent {
    return AsComponent(Shape_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TShape) GetNamePath() string {
    return Shape_GetNamePath(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TShape) Assign(Source IObject) {
    Shape_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TShape) ClassType() TClass {
    return Shape_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TShape) ClassName() string {
    return Shape_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TShape) InstanceSize() int32 {
    return Shape_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TShape) InheritsFrom(AClass TClass) bool {
    return Shape_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TShape) Equals(Obj IObject) bool {
    return Shape_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TShape) GetHashCode() int32 {
    return Shape_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TShape) ToString() string {
    return Shape_ToString(s.instance)
}

func (s *TShape) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Shape_AnchorToNeighbour(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (s *TShape) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Shape_AnchorParallel(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (s *TShape) AnchorHorizontalCenterTo(ASibling IControl) {
    Shape_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (s *TShape) AnchorVerticalCenterTo(ASibling IControl) {
    Shape_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TShape) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Shape_AnchorAsAlign(s.instance, ATheAlign , ASpace)
}

func (s *TShape) AnchorClient(ASpace int32) {
    Shape_AnchorClient(s.instance, ASpace)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (s *TShape) Align() TAlign {
    return Shape_GetAlign(s.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (s *TShape) SetAlign(value TAlign) {
    Shape_SetAlign(s.instance, value)
}

// 获取四个角位置的锚点。
func (s *TShape) Anchors() TAnchors {
    return Shape_GetAnchors(s.instance)
}

// 设置四个角位置的锚点。
func (s *TShape) SetAnchors(value TAnchors) {
    Shape_SetAnchors(s.instance, value)
}

// 获取画刷对象。
//
// Get Brush.
func (s *TShape) Brush() *TBrush {
    return AsBrush(Shape_GetBrush(s.instance))
}

// 设置画刷对象。
//
// Set Brush.
func (s *TShape) SetBrush(value *TBrush) {
    Shape_SetBrush(s.instance, CheckPtr(value))
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (s *TShape) DragCursor() TCursor {
    return Shape_GetDragCursor(s.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (s *TShape) SetDragCursor(value TCursor) {
    Shape_SetDragCursor(s.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (s *TShape) DragKind() TDragKind {
    return Shape_GetDragKind(s.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (s *TShape) SetDragKind(value TDragKind) {
    Shape_SetDragKind(s.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (s *TShape) DragMode() TDragMode {
    return Shape_GetDragMode(s.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (s *TShape) SetDragMode(value TDragMode) {
    Shape_SetDragMode(s.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (s *TShape) Enabled() bool {
    return Shape_GetEnabled(s.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (s *TShape) SetEnabled(value bool) {
    Shape_SetEnabled(s.instance, value)
}

// 获取约束控件大小。
func (s *TShape) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Shape_GetConstraints(s.instance))
}

// 设置约束控件大小。
func (s *TShape) SetConstraints(value *TSizeConstraints) {
    Shape_SetConstraints(s.instance, CheckPtr(value))
}

// 获取以父容器的ShowHint属性为准。
func (s *TShape) ParentShowHint() bool {
    return Shape_GetParentShowHint(s.instance)
}

// 设置以父容器的ShowHint属性为准。
func (s *TShape) SetParentShowHint(value bool) {
    Shape_SetParentShowHint(s.instance, value)
}

func (s *TShape) Pen() *TPen {
    return AsPen(Shape_GetPen(s.instance))
}

func (s *TShape) SetPen(value *TPen) {
    Shape_SetPen(s.instance, CheckPtr(value))
}

func (s *TShape) Shape() TShapeType {
    return Shape_GetShape(s.instance)
}

func (s *TShape) SetShape(value TShapeType) {
    Shape_SetShape(s.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (s *TShape) ShowHint() bool {
    return Shape_GetShowHint(s.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (s *TShape) SetShowHint(value bool) {
    Shape_SetShowHint(s.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (s *TShape) Visible() bool {
    return Shape_GetVisible(s.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (s *TShape) SetVisible(value bool) {
    Shape_SetVisible(s.instance, value)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (s *TShape) SetOnDragDrop(fn TDragDropEvent) {
    Shape_SetOnDragDrop(s.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (s *TShape) SetOnDragOver(fn TDragOverEvent) {
    Shape_SetOnDragOver(s.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (s *TShape) SetOnEndDrag(fn TEndDragEvent) {
    Shape_SetOnEndDrag(s.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (s *TShape) SetOnMouseDown(fn TMouseEvent) {
    Shape_SetOnMouseDown(s.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (s *TShape) SetOnMouseEnter(fn TNotifyEvent) {
    Shape_SetOnMouseEnter(s.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (s *TShape) SetOnMouseLeave(fn TNotifyEvent) {
    Shape_SetOnMouseLeave(s.instance, fn)
}

// 设置鼠标移动事件。
func (s *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
    Shape_SetOnMouseMove(s.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (s *TShape) SetOnMouseUp(fn TMouseEvent) {
    Shape_SetOnMouseUp(s.instance, fn)
}

func (s *TShape) Action() *TAction {
    return AsAction(Shape_GetAction(s.instance))
}

func (s *TShape) SetAction(value IComponent) {
    Shape_SetAction(s.instance, CheckPtr(value))
}

func (s *TShape) BiDiMode() TBiDiMode {
    return Shape_GetBiDiMode(s.instance)
}

func (s *TShape) SetBiDiMode(value TBiDiMode) {
    Shape_SetBiDiMode(s.instance, value)
}

func (s *TShape) BoundsRect() TRect {
    return Shape_GetBoundsRect(s.instance)
}

func (s *TShape) SetBoundsRect(value TRect) {
    Shape_SetBoundsRect(s.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (s *TShape) ClientHeight() int32 {
    return Shape_GetClientHeight(s.instance)
}

// 设置客户区高度。
//
// Set client height.
func (s *TShape) SetClientHeight(value int32) {
    Shape_SetClientHeight(s.instance, value)
}

func (s *TShape) ClientOrigin() TPoint {
    return Shape_GetClientOrigin(s.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (s *TShape) ClientRect() TRect {
    return Shape_GetClientRect(s.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (s *TShape) ClientWidth() int32 {
    return Shape_GetClientWidth(s.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (s *TShape) SetClientWidth(value int32) {
    Shape_SetClientWidth(s.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (s *TShape) ControlState() TControlState {
    return Shape_GetControlState(s.instance)
}

// 设置控件状态。
//
// Set control state.
func (s *TShape) SetControlState(value TControlState) {
    Shape_SetControlState(s.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (s *TShape) ControlStyle() TControlStyle {
    return Shape_GetControlStyle(s.instance)
}

// 设置控件样式。
//
// Set control style.
func (s *TShape) SetControlStyle(value TControlStyle) {
    Shape_SetControlStyle(s.instance, value)
}

func (s *TShape) Floating() bool {
    return Shape_GetFloating(s.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (s *TShape) Parent() *TWinControl {
    return AsWinControl(Shape_GetParent(s.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (s *TShape) SetParent(value IWinControl) {
    Shape_SetParent(s.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (s *TShape) Left() int32 {
    return Shape_GetLeft(s.instance)
}

// 设置左边位置。
//
// Set Left position.
func (s *TShape) SetLeft(value int32) {
    Shape_SetLeft(s.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (s *TShape) Top() int32 {
    return Shape_GetTop(s.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (s *TShape) SetTop(value int32) {
    Shape_SetTop(s.instance, value)
}

// 获取宽度。
//
// Get width.
func (s *TShape) Width() int32 {
    return Shape_GetWidth(s.instance)
}

// 设置宽度。
//
// Set width.
func (s *TShape) SetWidth(value int32) {
    Shape_SetWidth(s.instance, value)
}

// 获取高度。
//
// Get height.
func (s *TShape) Height() int32 {
    return Shape_GetHeight(s.instance)
}

// 设置高度。
//
// Set height.
func (s *TShape) SetHeight(value int32) {
    Shape_SetHeight(s.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (s *TShape) Cursor() TCursor {
    return Shape_GetCursor(s.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (s *TShape) SetCursor(value TCursor) {
    Shape_SetCursor(s.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (s *TShape) Hint() string {
    return Shape_GetHint(s.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (s *TShape) SetHint(value string) {
    Shape_SetHint(s.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TShape) ComponentCount() int32 {
    return Shape_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TShape) ComponentIndex() int32 {
    return Shape_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TShape) SetComponentIndex(value int32) {
    Shape_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TShape) Owner() *TComponent {
    return AsComponent(Shape_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TShape) Name() string {
    return Shape_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TShape) SetName(value string) {
    Shape_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TShape) Tag() int {
    return Shape_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TShape) SetTag(value int) {
    Shape_SetTag(s.instance, value)
}

// 获取左边锚点。
func (s *TShape) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Shape_GetAnchorSideLeft(s.instance))
}

// 设置左边锚点。
func (s *TShape) SetAnchorSideLeft(value *TAnchorSide) {
    Shape_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (s *TShape) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Shape_GetAnchorSideTop(s.instance))
}

// 设置顶边锚点。
func (s *TShape) SetAnchorSideTop(value *TAnchorSide) {
    Shape_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// 获取右边锚点。
func (s *TShape) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Shape_GetAnchorSideRight(s.instance))
}

// 设置右边锚点。
func (s *TShape) SetAnchorSideRight(value *TAnchorSide) {
    Shape_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// 获取底边锚点。
func (s *TShape) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Shape_GetAnchorSideBottom(s.instance))
}

// 设置底边锚点。
func (s *TShape) SetAnchorSideBottom(value *TAnchorSide) {
    Shape_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

// 获取边框间距。
func (s *TShape) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Shape_GetBorderSpacing(s.instance))
}

// 设置边框间距。
func (s *TShape) SetBorderSpacing(value *TControlBorderSpacing) {
    Shape_SetBorderSpacing(s.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TShape) Components(AIndex int32) *TComponent {
    return AsComponent(Shape_GetComponents(s.instance, AIndex))
}

// 获取锚侧面。
func (s *TShape) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Shape_GetAnchorSide(s.instance, AKind))
}

