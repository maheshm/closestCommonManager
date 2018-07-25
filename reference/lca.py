from pprint import pprint

class Node:
    def __init__(self, key):
        self.key = key
        self.children = []
        self.ancestor = None
        self.rank = 0
        self.color = 0
        self.parent = None

def printnodes():
    pprint(vars(p_u))
    pprint(vars(p_v))
    #pprint(vars(three))
    #pprint(vars(four))
    #pprint(vars(five))

def makeset(node):
    node.parent = node
    node.rank = 1

def union(nodex, nodey):
    xroot = find(nodex)
    yroot = find(nodey)
    if xroot.rank > yroot.rank:
        yroot.parent = xroot
    elif xroot.rank < yroot.rank:
        xroot.parent = yroot
    elif xroot.rank == yroot.rank:
        yroot.parent = xroot
        xroot.rank += 1

def find(node):
    if node.parent != node:
        node.parent = find(node.parent)
    return node.parent

def lca(u, found = False):
    #print "lca with %d" %(u.key)
    #printnodes()
    makeset(u)
    u.ancestor = u
    for w in u.children:
        res = lca(w, False)
        if res["found"]:
            return res
        #print "union u: %d, w: %d" %(u.key, w.key)
        union(u,w)
        #print "u.parent: %d, w.parent: %d, u.rank: %d, w.rank: %d" %(u.parent.key, w.parent.key, u.rank, w.rank)
        find(u).ancestor = u
        #print "find(u): %d, u_a: %d, w_a: %d, find(w).a: %d" %(find(u).key, u.ancestor.key, w.ancestor.key, find(w).ancestor.key)
    u.color = 1
    #print "p_u: %d, p_v: %d, u: %d, p_v.color: %d, p_u.color: %d, p_v_a: %d" %(p_u.key, p_v.key, u.key, p_v.color, p_u.color, (p_v.ancestor and p_v.ancestor.key) or 0)
    if (p_v.key == u.key) and (p_v.color == 1 and p_u.color == 1):
        #printnodes()
        #print find(p_u).key
        #print find(p_u).ancestor.key
        #print "\n"
        return {"a": find(p_u).ancestor, "found": True}
    elif (p_u.key == u.key) and (p_v.color == 1 and p_u.color == 1):
        #printnodes()
        #print find(p_v).key
        #print find(p_v).ancestor.key
        #print "hgf \n"
        return { "a": find(p_v).ancestor, "found": True}
    else:
        return {"found": False}

root = Node(1)
two = Node(2)
three = Node(3)
four = Node(4)
five = Node(5)
six = Node(6)
seven = Node(7)
eight = Node(8)

root.children.append(two)
root.children.append(three)
two.children.append(four)
two.children.append(five)
three.children.append(six)
three.children.append(seven)
three.children.append(eight)

p_u = two
p_v = eight

a = lca(root)
print a["a"].key
