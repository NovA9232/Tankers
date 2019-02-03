function debugVel(obj)
  lg.setColor({0, 1, 0})
  lg.line(obj.pos.x, obj.pos.y, obj.pos.x+obj.vel.x, obj.pos.y+obj.vel.y)
end
