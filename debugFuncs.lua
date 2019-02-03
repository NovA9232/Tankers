function debugVel(obj)
  lg.setColor({0, 1, 0})
  local x, y = obj:getMiddle()
  lg.line(x, y, obj.vel.y, obj.vel.x)
end
